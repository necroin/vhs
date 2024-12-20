package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
	"vhs/src/logger"
	"vhs/src/message"
	"vhs/src/network/connector"
	lan_listener "vhs/src/network/lan/listener"
	lan_observer "vhs/src/network/lan/observer"
	"vhs/src/network/server"
	"vhs/src/vhs/config"

	http_pprof "net/http/pprof"
)

type Application struct {
	config       *config.Config
	selfHostInfo *message.HostInfo
	hostsInfo    map[string]*message.HostInfo
	lanListener  *lan_listener.Listener
	lanObserver  *lan_observer.Observer
	connector    *connector.Connector
	server       *server.Server
	log          *logger.LogEntry
	services     map[string]string
	metrics      *Metrics
}

func New(config *config.Config, log *logger.LogEntry) (*Application, error) {
	log = log.WtihLabels("App")
	connector := connector.New(config)

	server := server.New(config, connector, log)

	app := &Application{
		config: config,
		selfHostInfo: &message.HostInfo{
			Url:      config.Url,
			Hostname: config.Hostname,
			Platform: config.Platform,
		},
		hostsInfo:   map[string]*message.HostInfo{},
		lanListener: lan_listener.New(config, log),
		lanObserver: lan_observer.New(config, log),
		connector:   connector,
		server:      server,
		log:         log,
		services:    map[string]string{},
		metrics:     NewMetrics(),
	}

	server.AddHandlerFunc("/", app.ServicesPageHandler, "GET")
	server.AddHandlerFunc("/services", app.ServicesHandler, "GET")
	server.AddHandlerFunc("/devices", app.DevicesHandler, "GET")
	server.AddHandlerFunc("/notify", app.NotifyHandler, "POST")
	server.AddHandler("/metrics", app.metrics.Registry.Handler(), "GET")

	server.AddHandlerFunc("/debug/pprof", http_pprof.Index, "GET")
	server.AddHandlerFunc("/debug/pprof/profile", http_pprof.Profile, "GET")

	log.Info("Read plugins")
	pluginsEntries, err := os.ReadDir(config.PluginsDir)
	if err != nil {
		return nil, log.NewError("failed read plugins directory: %s", err)
	}

	for _, entry := range pluginsEntries {
		pluginLog := log.WtihLabels("Plugins", entry.Name())
		pluginPath := path.Join(config.PluginsDir, entry.Name())
		pluginLog.Info("get plugin name")
		nameCallResult, err := app.CallPlugin(pluginPath, "name", []byte{}, CallOptions{})
		if err != nil {
			pluginLog.Error("failed get name of %s", pluginPath)
		}
		pluginLog.Info("plugin name is %s", nameCallResult)

		pluginLog.Info("get plugin services")
		servicesCallResult, err := app.CallPlugin(pluginPath, "services", []byte{}, CallOptions{})
		if err != nil {
			pluginLog.Error("failed get name of %s", pluginPath)
		}
		pluginLog.Info("plugin services: %s", servicesCallResult)

		services := map[string]message.ServiceInfo{}
		json.NewDecoder(bytes.NewBuffer(servicesCallResult)).Decode(&services)

		for serviceName, serviceInfo := range services {
			serviceEndpoint := "/" + path.Join(string(nameCallResult), serviceInfo.Endpoint)
			pluginLog.Info("add service %s by path %s", serviceName, serviceEndpoint)

			suppressLogs := serviceInfo.SuppressLogs
			server.AddHandlerFunc(
				serviceEndpoint,
				func(responseWriter http.ResponseWriter, request *http.Request) {
					serviceLog := pluginLog.WtihLabels(serviceName)
					if !suppressLogs {
						serviceLog.Verbose("called")
					}

					data, err := io.ReadAll(request.Body)
					if err != nil {
						serviceLog.Error("failed read data: %s", err.Error())
						return
					}

					if _, err = app.CallPlugin(
						pluginPath,
						serviceName,
						data,
						CallOptions{
							CmdOut:       responseWriter,
							SuppressLogs: suppressLogs,
						},
					); err != nil {
						serviceLog.Error(err.Error())
						return
					}
				},
				serviceInfo.Methods...,
			)

			if serviceName == "page" {
				app.services[string(nameCallResult)] = serviceEndpoint
			}
		}
	}

	return app, nil
}

func (app *Application) Start() error {
	wg := sync.WaitGroup{}

	app.log.Info("Start server")
	wg.Add(1)
	go func() {
		defer wg.Done()
		app.server.Start()
	}()

	if err := app.server.WaitStart(); err != nil {
		return err
	}

	app.log.Info("Start lan listener")
	go func() {
		addrs, _ := app.lanListener.Start()

		for {
			notifyAddr := <-addrs
			go func() {
				if err := app.Notify(notifyAddr); err != nil {
					app.log.Error(err.Error())
				}
			}()
		}
	}()

	go func() {
		time.Sleep(config.InstanceRemoveTimeout)
		currentTime := time.Now().UnixNano()
		currentTimeSeconds := time.Duration(currentTime).Seconds()

		deleteHosts := []string{}

		for hostname, host := range app.hostsInfo {
			timestampSeconds := time.Duration(host.Timestamp).Seconds()
			if currentTimeSeconds-timestampSeconds > config.InstanceRemoveTimeout.Seconds() {
				deleteHosts = append(deleteHosts, hostname)
			}
		}

		for _, hostname := range deleteHosts {
			delete(app.hostsInfo, hostname)
		}
	}()

	app.log.Info("Start lan observer")
	app.lanObserver.Start()

	app.log.Info("Platform is %s", app.config.Platform)
	app.log.Info("Server started on %s", app.config.Url)

	wg.Wait()

	return nil
}

func (app *Application) Notify(url string) error {
	app.log.WtihLabels("Notify").Debug(url)
	message := message.HostInfo{
		Url:       app.config.Url,
		Hostname:  app.config.Hostname,
		Platform:  app.config.Platform,
		Services:  app.services,
		Timestamp: time.Now().UnixNano(),
	}

	_, err := app.connector.SendPostRequest(
		url+"/notify",
		message,
	)

	return err
}

package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vhs/src/logger"
	"vhs/src/network/connector"
	"vhs/src/vhs/config"

	"github.com/gorilla/mux"
)

const (
	ServerStatusEndpoint        = "/status"
	ServerWaitStartRepeatCount  = 10
	ServerWaitStartSleepSeconds = 1
	ServerStatusResponse        = "OK"
)

type Server struct {
	config    *config.Config
	router    *mux.Router
	instance  *http.Server
	connector *connector.Connector
	log       *logger.LogEntry
}

func New(config *config.Config, connector *connector.Connector, log *logger.LogEntry) *Server {
	router := mux.NewRouter()

	instance := &http.Server{
		Addr:    config.Url,
		Handler: router,
	}

	server := &Server{
		config:    config,
		router:    router,
		instance:  instance,
		connector: connector,
		log:       log.WtihLabels("Server"),
	}

	server.AddHandlerFunc(ServerStatusEndpoint, func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Write([]byte(ServerStatusResponse))
	}, "GET")

	return server
}

func (server *Server) Start() {
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		server.instance.Shutdown(ctx)
	}()

	server.instance.ListenAndServe()
}

func (server *Server) AddHandler(path string, handler http.Handler, methods ...string) {
	server.router.Handle(path, handler).Methods(methods...)
}

func (server *Server) AddHandlerFunc(path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	server.router.HandleFunc(path, handler).Methods(methods...)
}

func (server *Server) WaitStart() error {
	log := server.log.WtihLabels("WaitStart")
	for i := 0; i < ServerWaitStartRepeatCount; i++ {
		response, err := server.connector.SendRequest(server.config.Url+ServerStatusEndpoint, []byte(""), http.MethodGet)
		if err != nil {
			log.Error("failed send request: %s", err)
			time.Sleep(ServerWaitStartSleepSeconds * time.Second)
			continue
		}
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error("failed read response data: %s", err)
			time.Sleep(ServerWaitStartSleepSeconds * time.Second)
			continue
		}
		if string(data) == ServerStatusResponse {
			return nil
		}
	}
	return fmt.Errorf("[Server] [WaitStart] [Error] failed get server status")
}

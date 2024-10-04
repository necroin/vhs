package lan_observer

import (
	"fmt"
	"net"
	"strings"
	"time"
	"vhs/src/logger"
	"vhs/src/vhs/config"
)

type Observer struct {
	config *config.Config
	log    *logger.LogEntry
}

func New(config *config.Config, log *logger.LogEntry) *Observer {
	return &Observer{
		config: config,
		log:    log.WtihLabels("Lan Observer"),
	}
}

func (observer *Observer) Start() error {
	lanNetwork3 := "0"

	if !strings.Contains(observer.config.Url, "localhost") {
		lanNetworkParts := strings.Split(observer.config.Url, ".")
		lanNetwork3 = lanNetworkParts[2]
	}

	local, err := net.ResolveUDPAddr("udp4", observer.config.Url)
	if err != nil {
		return err
	}

	remote, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("192.168.%s.255:%s", lanNetwork3, observer.config.TopologyPort))
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp4", local, remote)
	if err != nil {
		return err
	}

	go func() {
		for {
			_, err = conn.Write([]byte("vhs observe"))
			if err != nil {
				observer.log.Error("%s", err)
			}
			time.Sleep(observer.config.RequestTimeout)
		}
	}()

	return nil
}

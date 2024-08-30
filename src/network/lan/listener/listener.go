package lan_listener

import (
	"fmt"
	"net"
	"vhs/src/logger"
	"vhs/src/vhs/config"
)

type Listener struct {
	config *config.Config
	log    *logger.LogEntry
}

func New(config *config.Config, log *logger.LogEntry) *Listener {
	return &Listener{
		config: config,
		log:    log.WtihLabels("Lan Listener"),
	}
}

func (listener *Listener) Start() (chan string, error) {
	result := make(chan string)

	udpAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%s", listener.config.ListenPort))
	if err != nil {
		return nil, err
	}

	netListener, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		return nil, err
	}
	go func() {
		buf := make([]byte, 1024)
		for {
			_, addr, err := netListener.ReadFromUDP(buf)
			if err != nil {
				panic(err)
			}
			listener.log.Debug("%s", addr.String())
			result <- addr.String()
		}
	}()
	return result, nil
}

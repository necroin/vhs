package app

import (
	"encoding/json"
	"net/http"
	"vhs/src/message"
)

func (app *Application) NotifyHandler(responseWriter http.ResponseWriter, request *http.Request) {
	log := app.log.WtihLabels("Notify Handler")
	log.Verbose("handler called")
	hostInfo := &message.HostInfo{}

	if err := json.NewDecoder(request.Body).Decode(hostInfo); err != nil {
		log.Error("failed decode host info: %s", err)
	}

	log.Debug("host info: %v", hostInfo)
	app.hostsInfo[hostInfo.Hostname] = hostInfo
}

func (app *Application) ServicesPageHandler(responseWriter http.ResponseWriter, request *http.Request) {

}

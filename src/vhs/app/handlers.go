package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"vhs/src/message"
	plugins_core "vhs/src/plugins/core"
)

const (
	BaseStylePath          = "assets/web_interface/style.css"
	ServicesPageHtmlPath   = "assets/web_interface/services/services.html"
	ServicesPageStylePath  = "assets/web_interface/services/services.css"
	ServicesPageScriptPath = "assets/web_interface/services/services.js"
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
	pageHtmlData, err := os.ReadFile(ServicesPageHtmlPath)
	if err != nil {
		responseWriter.Write([]byte(err.Error()))
		return
	}

	baseStyleData, err := os.ReadFile(BaseStylePath)
	if err != nil {
		responseWriter.Write([]byte(err.Error()))
		return
	}

	pageStyleData, err := os.ReadFile(ServicesPageStylePath)
	if err != nil {
		responseWriter.Write([]byte(err.Error()))
		return
	}

	pageScriptData, err := os.ReadFile(ServicesPageScriptPath)
	if err != nil {
		responseWriter.Write([]byte(err.Error()))
		return
	}

	pageInfo := plugins_core.PageInfo{
		BaseStyle: fmt.Sprintf("<style>%s</style>", baseStyleData),
		Style:     fmt.Sprintf("<style>%s</style>", pageStyleData),
		Script:    fmt.Sprintf("<script>%s</script>", pageScriptData),
	}

	pageTemplate, err := template.New("page").Parse(string(pageHtmlData))
	if err != nil {
		responseWriter.Write([]byte(err.Error()))
		return
	}

	if err := pageTemplate.Execute(responseWriter, pageInfo); err != nil {
		responseWriter.Write([]byte(err.Error()))
		return
	}
}

func (app *Application) ServicesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if err := json.NewEncoder(responseWriter).Encode(app.services); err != nil {
		responseWriter.Write([]byte(err.Error()))
	}
}

func (app *Application) DevicesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if err := json.NewEncoder(responseWriter).Encode(app.hostsInfo); err != nil {
		responseWriter.Write([]byte(err.Error()))
	}
}

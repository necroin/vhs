package main

import (
	"fmt"
	"net/http"
	"os"
	"vhs/src/logger"
	"vhs/src/vhs/app"
	"vhs/src/vhs/config"
)

func NotifyHandler(responseWriter http.ResponseWriter, request *http.Request) {

}

func main() {
	config, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger := logger.NewLogger(os.Stdout, logger.GetLogLevel(config.Log.Level))
	log := logger.WtihLabels("main")

	log.Info("Create aplication")
	app, err := app.New(config, logger.WtihLabels())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Info("Start aplication")
	if err := app.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"path"
	"sync"
	"time"
	plugins_core "vhs/src/plugins/core"
	"vhs/src/utils"
)

type CallOptions struct {
	CmdOut       io.Writer
	SuppressLogs bool
}

func (app *Application) CallPlugin(pluginPath string, serviceName string, data []byte, options CallOptions) ([]byte, error) {
	clock := utils.NewClock(func(delta time.Duration) {
		app.metrics.Storage.CallTimeHistogramVector.WithLabelValues(pluginPath, serviceName).Observe(float64(delta.Milliseconds()))
	})
	defer clock.Stop()

	log := app.log.WtihLabels("Call Plugin", serviceName)
	execTool, execArgs := utils.GetProcessRunCommand(app.config.Platform, path.Clean(pluginPath))
	execArgs = append(execArgs, fmt.Sprintf("-call=%s", serviceName))
	cmd := exec.Command(execTool, execArgs...)
	if !options.SuppressLogs {
		log.Debug("cmd command: %s", cmd.Args)
	}

	cmdIn := &bytes.Buffer{}
	cmdOut := &bytes.Buffer{}

	input := &plugins_core.ServiceInput{
		ClusterInfo: plugins_core.ClusterInfo{
			Self:  app.selfHostInfo,
			Hosts: app.hostsInfo,
		},
		Data: data,
	}
	if !options.SuppressLogs {
		log.Debug("input: %v", input)
	}
	if err := json.NewEncoder(cmdIn).Encode(input); err != nil {
		return nil, log.NewError("failed encode service input: %s", err)
	}

	cmd.Stdin = cmdIn
	cmd.Stdout = cmdOut
	cmd.Stderr = cmdOut

	if options.CmdOut != nil {
		cmd.Stdout = options.CmdOut
		cmd.Stderr = options.CmdOut
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if !options.SuppressLogs {
			log.Verbose("Run process")
		}
		if err := cmd.Run(); err != nil {
			log.Error("failed start process: %s", err)
		}
	}()
	wg.Wait()

	return cmdOut.Bytes(), nil
}

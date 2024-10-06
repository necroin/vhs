package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"path"
	plugins_core "vhs/src/plugins/core"
	"vhs/src/utils"
)

func (app *Application) CallPlugin(pluginPath string, serviceName string, data []byte, suppressLogs bool) ([]byte, error) {
	log := app.log.WtihLabels("Call Plugin", serviceName)
	execTool, execArgs := utils.GetProcessRunCommand(app.config.Platform, path.Clean(pluginPath))
	execArgs = append(execArgs, fmt.Sprintf("-call=%s", serviceName))
	cmd := exec.Command(execTool, execArgs...)
	if !suppressLogs {
		log.Debug("cmd command: %s", cmd.Args)
	}

	myPipeReader, handlerPipeWriter := io.Pipe()
	defer myPipeReader.Close()
	defer handlerPipeWriter.Close()

	cmdIn := &bytes.Buffer{}
	input := &plugins_core.ServiceInput{
		ClusterInfo: plugins_core.ClusterInfo{
			Self:  app.selfHostInfo,
			Hosts: app.hostsInfo,
		},
		Data: data,
	}
	if !suppressLogs {
		log.Debug("input: %v", input)
	}
	if err := json.NewEncoder(cmdIn).Encode(input); err != nil {
		return nil, log.NewError("failed encode service input: %s", err)
	}

	cmd.Stdin = cmdIn
	cmd.Stdout = handlerPipeWriter
	cmd.Stderr = handlerPipeWriter

	go func() {
		defer handlerPipeWriter.Close()
		if !suppressLogs {
			log.Verbose("Run process")
		}
		if err := cmd.Run(); err != nil {
			log.Error("failed start process: %s", err)
		}
	}()

	cmdOut := &bytes.Buffer{}
	if _, err := io.Copy(cmdOut, myPipeReader); err != nil {
		return nil, log.NewError("failed copy process result: %s", err)
	}

	return cmdOut.Bytes(), nil
}

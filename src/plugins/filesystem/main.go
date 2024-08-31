package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	plugins_core "vhs/src/plugins/core"
	filesystem_handlers "vhs/src/plugins/filesystem/handlers"
)

var (
	call = map[string]func(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error{
		filesystem_handlers.NameServiceName:           filesystem_handlers.NameHandler,
		filesystem_handlers.ServicesServiceName:       filesystem_handlers.ServicesHandler,
		filesystem_handlers.PageServiceName:           filesystem_handlers.PageHandler,
		filesystem_handlers.DevicesServiceName:        filesystem_handlers.DevicesHandler,
		filesystem_handlers.FilesystemSelfServiceName: filesystem_handlers.FilesystemSelfHandler,
		filesystem_handlers.FilesystemAllServiceName:  filesystem_handlers.FilesystemAllHandler,
		filesystem_handlers.CreateServiceName:         filesystem_handlers.CreateHandler,
		filesystem_handlers.DeleteServiceName:         filesystem_handlers.DeleteHandler,
		filesystem_handlers.SelectServiceName:         filesystem_handlers.SelectHandler,
		filesystem_handlers.CopyServiceName:           filesystem_handlers.CopyHandler,
		filesystem_handlers.MoveServiceName:           filesystem_handlers.MoveHandler,
		filesystem_handlers.RenameServiceName:         filesystem_handlers.RenameHandler,
	}
)

func main() {
	flags := plugins_core.ParseFlags()

	serviceInput := &plugins_core.ServiceInput{}
	json.NewDecoder(os.Stdin).Decode(serviceInput)

	callFunc, ok := call[flags.Call]
	if !ok {
		fmt.Fprintf(os.Stderr, "[filesystem] wrong service: %s", flags.Call)
		return
	}

	if err := callFunc(&serviceInput.ClusterInfo, os.Stdout, serviceInput.Data); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}

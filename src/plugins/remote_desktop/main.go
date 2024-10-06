package main

import (
	"io"
	plugins_core "vhs/src/plugins/core"
	remote_desktop_handlers "vhs/src/plugins/remote_desktop/handlers"
)

var (
	call = map[string]func(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error{
		remote_desktop_handlers.NameServiceName:          remote_desktop_handlers.NameHandler,
		remote_desktop_handlers.ServicesServiceName:      remote_desktop_handlers.ServicesHandler,
		remote_desktop_handlers.PageServiceName:          remote_desktop_handlers.PageHandler,
		remote_desktop_handlers.ImageServiceName:         remote_desktop_handlers.ImageHandler,
		remote_desktop_handlers.StreamServiceName:        remote_desktop_handlers.StreamHandler,
		remote_desktop_handlers.InputMouseServiceName:    remote_desktop_handlers.MouseEventHandler,
		remote_desktop_handlers.InputKeyboardServiceName: remote_desktop_handlers.KeyboardEventHandler,
	}
)

func main() {
	plugins_core.MainPipeline(call)
}

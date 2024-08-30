package filesystem_handlers

import (
	"encoding/json"
	"io"
	"vhs/src/message"
	plugins_core "vhs/src/plugins/core"
)

func ServicesHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	result := map[string]message.ServiceInfo{
		PageServiceName: {
			Endpoint: "/explorer",
		},
		DevicesServiceName: {
			Endpoint: "/filesystem/devices",
		},
		FilesystemSelfServiceName: {
			Endpoint: "/filesystem/self",
		},
		FilesystemAllServiceName: {
			Endpoint: "/filesystem/all",
		},
		CreateServiceName: {
			Endpoint: "/filesystem/create",
		},
	}
	return json.NewEncoder(out).Encode(result)
}

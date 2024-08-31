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
			Methods:  []string{"GET"},
		},
		DevicesServiceName: {
			Endpoint: "/devices",
			Methods:  []string{"GET"},
		},
		FilesystemSelfServiceName: {
			Endpoint: "/self",
			Methods:  []string{"POST", "GET"},
		},
		FilesystemAllServiceName: {
			Endpoint: "/all",
			Methods:  []string{"POST", "GET"},
		},
		CreateServiceName: {
			Endpoint: "/create",
			Methods:  []string{"POST"},
		},
		DeleteServiceName: {
			Endpoint: "/delete",
			Methods:  []string{"POST"},
		},
		SelectServiceName: {
			Endpoint: "/select",
			Methods:  []string{"POST"},
		},
		CopyServiceName: {
			Endpoint: "/copy",
			Methods:  []string{"POST"},
		},
		MoveServiceName: {
			Endpoint: "/move",
			Methods:  []string{"POST"},
		},
		RenameServiceName: {
			Endpoint: "/rename",
			Methods:  []string{"POST"},
		},
	}
	return json.NewEncoder(out).Encode(result)
}

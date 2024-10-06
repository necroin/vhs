package remote_desktop_handlers

import (
	"encoding/json"
	"io"
	"vhs/src/message"
	plugins_core "vhs/src/plugins/core"
)

func ServicesHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	result := map[string]message.ServiceInfo{
		PageServiceName: {
			Endpoint: "/hosts",
			Methods:  []string{"GET"},
		},
		ImageServiceName: {
			Endpoint:     "/image",
			Methods:      []string{"GET"},
			SuppressLogs: true,
		},
		StreamServiceName: {
			Endpoint: "/stream",
			Methods:  []string{"GET"},
		},
		InputMouseServiceName: {
			Endpoint: "/input/mouse",
			Methods:  []string{"POST"},
		},
		InputKeyboardServiceName: {
			Endpoint: "/input/keyboard",
			Methods:  []string{"POST"},
		},
	}
	return json.NewEncoder(out).Encode(result)
}

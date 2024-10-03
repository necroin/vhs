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
	}
	return json.NewEncoder(out).Encode(result)
}

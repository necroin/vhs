package filesystem_handlers

import (
	"encoding/json"
	"io"
	plugins_core "vhs/src/plugins/core"
)

func DevicesHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	return json.NewEncoder(out).Encode(clusterInfo.Hosts)
}

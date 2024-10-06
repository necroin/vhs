package remote_desktop_handlers

import (
	"io"
	plugins_core "vhs/src/plugins/core"
)

func StreamHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	return plugins_core.PageHandler(
		"assets/web_interface/plugins/remote_desktop/stream/stream.html",
		"assets/web_interface/plugins/remote_desktop/stream/stream.css",
		"assets/web_interface/plugins/remote_desktop/stream/stream.js",
		clusterInfo,
		out,
		data,
	)
}

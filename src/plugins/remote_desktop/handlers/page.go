package remote_desktop_handlers

import (
	"io"
	plugins_core "vhs/src/plugins/core"
)

func PageHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	return plugins_core.PageHandler(
		"assets/web_interface/plugins/remote_desktop/hosts/hosts.html",
		"assets/web_interface/plugins/remote_desktop/hosts/hosts.css",
		"assets/web_interface/plugins/remote_desktop/hosts/hosts.js",
		clusterInfo,
		out,
		data,
	)
}

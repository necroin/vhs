package filesystem_handlers

import (
	"io"
	plugins_core "vhs/src/plugins/core"
)

func PageHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	return plugins_core.PageHandler(
		"assets/web_interface/plugins/filesystem/explorer/explorer.html",
		"assets/web_interface/plugins/filesystem/explorer/explorer.css",
		"assets/web_interface/plugins/filesystem/explorer/explorer.js",
		clusterInfo,
		out,
		data,
	)
}

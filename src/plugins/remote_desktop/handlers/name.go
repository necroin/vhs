package remote_desktop_handlers

import (
	"fmt"
	"io"
	plugins_core "vhs/src/plugins/core"
)

func NameHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	fmt.Fprint(out, "remote_desktop")
	return nil
}

package filesystem_handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"path"
	plugins_core "vhs/src/plugins/core"
	filesystem_utils "vhs/src/plugins/filesystem/utils"
)

func SelectHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	operation := &SelectOperation{}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(operation); err != nil {
		return err
	}
	operation.Path = path.Clean(operation.Path)

	if err := filesystem_utils.Compress(operation.Path, out); err != nil {
		return err
	}

	return nil
}

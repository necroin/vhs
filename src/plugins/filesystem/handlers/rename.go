package filesystem_handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"path"
	plugins_core "vhs/src/plugins/core"
	filesystem_utils "vhs/src/plugins/filesystem/utils"
)

func RenameHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	operation := &RenameOperation{}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(operation); err != nil {
		return err
	}
	operation.SrcPath = path.Clean(operation.SrcPath)
	operation.DstPath = path.Clean(operation.DstPath)

	if err := filesystem_utils.Rename(operation.SrcPath, operation.DstPath); err != nil {
		return err
	}

	return nil
}

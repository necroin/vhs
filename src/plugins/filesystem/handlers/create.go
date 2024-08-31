package filesystem_handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"path"
	plugins_core "vhs/src/plugins/core"
	filesystem_utils "vhs/src/plugins/filesystem/utils"
)

var (
	createHandlers = map[string]func(string) error{
		"directory": filesystem_utils.CreateNewDirectory,
		"file":      filesystem_utils.CreateNewFile,
	}
)

func CreateHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	operation := &CreateOperation{}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(operation); err != nil {
		return err
	}
	operation.Path = path.Clean(operation.Path)

	handler := createHandlers[operation.Type]
	if err := handler(operation.Path); err != nil {
		return err
	}

	return nil
}

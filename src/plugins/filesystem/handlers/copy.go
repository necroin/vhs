package filesystem_handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"path"
	"vhs/src/network/http"
	plugins_core "vhs/src/plugins/core"
	filesystem_utils "vhs/src/plugins/filesystem/utils"
)

func CopyHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	operation := &CopyOperation{}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(operation); err != nil {
		return err
	}
	operation.SrcPath = path.Clean(operation.SrcPath)
	operation.DstPath = path.Clean(operation.DstPath)

	selectOperation := &SelectOperation{
		Path: operation.SrcPath,
	}
	selectOperationData, err := json.Marshal(selectOperation)
	if err != nil {
		return err
	}

	response, err := http.SendPostRequest(operation.SrcUrl+"/filesystem/select", selectOperationData)
	if err != nil {
		return err
	}

	if err := filesystem_utils.Decompress(response.Body, operation.DstPath); err != nil {
		return err
	}

	return nil
}

package filesystem_handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"path"
	"vhs/src/network/http"
	plugins_core "vhs/src/plugins/core"
	filesystem_utils "vhs/src/plugins/filesystem/utils"
)

func MoveHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	operation := &MoveOperation{}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(operation); err != nil {
		return err
	}
	operation.SrcPath = path.Clean(operation.SrcPath)
	operation.DstPath = path.Clean(operation.DstPath)

	if clusterInfo.Self.Url == operation.SrcUrl {
		return filesystem_utils.Rename(operation.SrcPath, operation.DstPath)
	}

	if err := CopyHandler(clusterInfo, out, data); err != nil {
		return err
	}

	deleteOperation := &DeleteOperation{
		Path: operation.SrcPath,
	}
	deleteOperationData, err := json.Marshal(deleteOperation)
	if err != nil {
		return err
	}

	deleteResponse, err := http.SendPostRequest(operation.SrcUrl+"/filesystem/delete", deleteOperationData)
	if err != nil {
		return err
	}

	deleteResponseData, err := io.ReadAll(deleteResponse.Body)
	if err != nil {
		return err
	}

	if len(deleteResponseData) > 0 {
		return errors.New(string(deleteResponseData))
	}

	return nil
}

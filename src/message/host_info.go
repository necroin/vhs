package message

import (
	"bytes"
	"encoding/json"
)

type HostInfo struct {
	Url       string `json:"url"`
	Hostname  string `json:"hostname"`
	Platform  string `json:"platform"`
	Timestamp int64  `json:"timestamp"`
}

func (hostInfo *HostInfo) String() string {
	result := &bytes.Buffer{}
	json.NewEncoder(result).Encode(hostInfo)
	return result.String()
}

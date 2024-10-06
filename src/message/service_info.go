package message

import (
	"bytes"
	"encoding/json"
)

type ServiceInfo struct {
	Endpoint     string   `json:"endpoint"`
	Methods      []string `json:"methods"`
	SuppressLogs bool     `json:"suppress_logs"`
}

func (serviceInfo *ServiceInfo) String() string {
	result := &bytes.Buffer{}
	json.NewEncoder(result).Encode(serviceInfo)
	return result.String()
}

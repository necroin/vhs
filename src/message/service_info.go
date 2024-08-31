package message

type ServiceInfo struct {
	Endpoint string   `json:"endpoint"`
	Methods  []string `json:"methods"`
}

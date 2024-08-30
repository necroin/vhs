package message

type HostInfo struct {
	Url       string `json:"url"`
	Hostname  string `json:"hostname"`
	Platform  string `json:"platform"`
	Timestamp int64  `json:"timestamp"`
}

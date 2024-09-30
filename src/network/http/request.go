package http

import (
	"bytes"
	"net/http"
	"time"
)

func SendRequest(url string, data []byte, method string, timeout time.Duration) (*http.Response, error) {
	proto := "http://"

	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(
		method,
		proto+url,
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

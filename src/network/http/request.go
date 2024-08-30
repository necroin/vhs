package http

import (
	"bytes"
	"net/http"
)

func SendRequest(url string, data []byte, method string) (*http.Response, error) {
	proto := "http://"

	client := http.Client{}

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

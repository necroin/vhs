package http

import "net/http"

func SendPostRequest(url string, data []byte) (*http.Response, error) {
	return SendRequest(url, data, http.MethodPost)
}

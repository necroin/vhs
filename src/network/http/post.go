package http

import (
	"net/http"
	"time"
)

func SendPostRequest(url string, data []byte) (*http.Response, error) {
	return SendRequest(url, data, http.MethodPost, time.Second*5)
}

package connector

import (
	"bytes"
	"encoding/json"
	"net/http"
	"vhs/src/vhs/config"
)

type Connector struct {
	config *config.Config
}

func New(config *config.Config) *Connector {
	return &Connector{
		config: config,
	}
}

func (connector *Connector) SendPostRequest(url string, data any) (*http.Response, error) {
	return connector.SendRequestWithDataEncode(url, data, http.MethodPost)
}

func (connector *Connector) SendGetRequest(url string, data []byte, result any) error {
	response, err := connector.SendRequest(url, data, http.MethodGet)
	if err != nil {
		return err
	}

	return json.NewDecoder(response.Body).Decode(result)
}

func (connector *Connector) SendRequestWithDataEncode(url string, data any, method string) (*http.Response, error) {
	encodedMessage, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return connector.SendRequest(url, encodedMessage, method)
}

func (connector *Connector) SendRequest(url string, data []byte, method string) (*http.Response, error) {
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

package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseUrl    string
	HttpClient *http.Client
}

type HTTPHeader struct {
	Key   string
	Value string
}

type HTTPResponse struct {
	Body       []byte
	StatusCode int
	Status     string
}

func New(baseUrl string,
	httpClient *http.Client) Client {
	if httpClient != nil {
		return Client{BaseUrl: baseUrl, HttpClient: httpClient}
	}
	return Client{BaseUrl: baseUrl, HttpClient: &http.Client{Timeout: time.Duration(1) * time.Second}}
}

func (client *Client) Request(method string, path string, body interface{}, headers ...HTTPHeader) (*HTTPResponse, error) {
	var requestBody io.Reader
	if body != nil {
		data, _ := json.Marshal(body)
		requestBody = bytes.NewBuffer(data)
	}

	url := fmt.Sprintf("%s%s", client.BaseUrl, path)

	request, requestErr := http.NewRequest(method, url, requestBody)
	if requestErr != nil {
		return nil, requestErr
	}

	request.Header.Add("Accept", `application/json`)

	for _, header := range headers {
		request.Header.Add(header.Key, header.Value)
	}

	response, responseErr := client.HttpClient.Do(request)
	if responseErr != nil {
		return nil, responseErr
	}

	defer func() {
		_ = response.Body.Close()
	}()

	data, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	return &HTTPResponse{Body: data, StatusCode: response.StatusCode, Status: response.Status}, nil
}

func (client *Client) Get(path string, headers ...HTTPHeader) (*HTTPResponse, error) {
	return client.Request("GET", path, nil, headers...)
}

func (client *Client) Post(path string, body interface{}, headers ...HTTPHeader) (*HTTPResponse, error) {
	return client.Request("POST", path, body, headers...)
}

func (client *Client) Put(path string, body interface{}, headers ...HTTPHeader) (*HTTPResponse, error) {
	return client.Request("PUT", path, body, headers...)
}

func (client *Client) Patch(path string, body interface{}, headers ...HTTPHeader) (*HTTPResponse, error) {
	return client.Request("Patch", path, body, headers...)
}

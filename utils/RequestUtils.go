package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestHeader struct {
	headers map[string]string
}

var (
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
)

func NewRequest(method string, url string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req, nil

}

func GetWithAccessToken(url string, accessToken string) (int, []byte, error) {
	req, err := NewRequest(
		"GET",
		url,
		NewRequestHeaders().Header("Authorization", fmt.Sprintf("Bearer %s", accessToken)).Build(),
		nil)

	if err != nil {
		return -1, nil, err
	}

	return makeRequest(req)
}

func PostJsonModel(url string, authorization string, obj interface{}) (int, []byte, error) {
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return -1, nil, err
	}
	req, err := NewRequest(
		"POST",
		url,
		NewRequestHeaders().
			Header(HeaderAuthorization, authorization).
			Header(HeaderContentType, "application/json").
			Build(),
		bytes.NewBuffer([]byte(jsonStr)))

	if err != nil {
		return -1, nil, err
	}

	return makeRequest(req)
}

func makeRequest(req *http.Request) (int, []byte, error) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}
	return resp.StatusCode, data, nil
}

func (h *RequestHeader) Header(key string, value string) *RequestHeader {
	h.headers[key] = value
	return h
}

func (h *RequestHeader) Build() map[string]string {
	return h.headers
}

func NewRequestHeaders() *RequestHeader {
	return &RequestHeader{
		headers: make(map[string]string),
	}
}

package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

func (c *client) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	return client.Do(request)
}

func (c *client) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}

func (c *client) getRequestHeaders(requestHeaders http.Header) http.Header {
	headers := make(http.Header)

	// Add custom headers to the request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			headers.Set(header, value[0])
		}
	}

	// Add common headers to the request
	for header, value := range c.Headers {
		if len(value) > 0 {
			headers.Set(header, value[0])
		}
	}

	return headers
}

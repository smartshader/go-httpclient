package gohttp

import (
	"errors"
	"net/http"
)

func (c *client) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	fullHeaders := c.getRequestHeaders(headers)
	request.Header = fullHeaders

	return client.Do(request)
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

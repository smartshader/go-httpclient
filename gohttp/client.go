package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client *http.Client

	maxIdleConnections int
	responseTimeout    time.Duration
	connectionTimeout  time.Duration
	disableTimeouts    bool

	Headers http.Header
}

func New() HttpClient {
	return &httpClient{}
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	SetResponseTimeout(timeout time.Duration)
	SetConnectionTimeout(timeout time.Duration)
	SetMaxIdleConnections(max int)
	DisableTimeouts(disable bool)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) DisableTimeouts(disable bool) {
	c.disableTimeouts = disable
}

func (c *httpClient) SetResponseTimeout(timeout time.Duration) {
	c.responseTimeout = timeout
}

func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

func (c *httpClient) SetMaxIdleConnections(max int) {
	c.maxIdleConnections = max
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

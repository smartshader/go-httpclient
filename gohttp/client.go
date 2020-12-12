package gohttp

import "net/http"

type client struct{}

func New() Client {
	return &client{}
}

type Client interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *client) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *client) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *client) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *client) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *client) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	responseTimeout    time.Duration
	connectionTimeout  time.Duration
	disableTimeouts    bool
}

type ClientBuilder interface {
	Build() Client

	SetHeaders(headers http.Header) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(max int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
}

func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

func (c *clientBuilder) Build() Client {
	return &httpClient{
		builder: c,
	}
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(max int) ClientBuilder {
	c.maxIdleConnections = max
	return c
}

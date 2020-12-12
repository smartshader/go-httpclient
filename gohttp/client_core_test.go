package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	httpClient := &client{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	httpClient.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	fullHeaders := httpClient.getRequestHeaders(requestHeaders)

	if len(fullHeaders) != 3 {
		t.Error("we expect 3 headers")
	}

	if fullHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}

	if fullHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid request content type received")
	}

	if fullHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid request user agent received")
	}
}

func TestGetRequestBody(t *testing.T) {
	httpClient := &client{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := httpClient.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when parsing a nil body")
		}

		if body != nil {
			t.Error("no body expected when parsing a nil body")
		}
	})

	t.Run("bodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := httpClient.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling a slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("bodyWithXml", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := httpClient.getRequestBody("application/xml", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling a slice as xml")
		}

		if string(body) != "<string>one</string><string>two</string>" {
			t.Error("invalid xml body obtained")
		}
	})

	t.Run("bodyWithJsonAsDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := httpClient.getRequestBody("None", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling a slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})
}

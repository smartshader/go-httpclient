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
}

package main

import (
	"fmt"
	"github.com/smartshader/go-httpclient/gohttp"
	"net/http"
)

var (
	httpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		SetHeaders(commonHeaders).Build()

	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode())
	fmt.Println(response.Status())
	fmt.Println(response.String())
}

func createUser(user User) {
	response, err := httpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response.Bytes()))
}

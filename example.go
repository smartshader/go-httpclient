package main

import (
	"fmt"
	"github.com/smartshader/go-httpclient/gohttp"
	"io/ioutil"
	"net/http"
)

var (
	httpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client.SetHeaders(commonHeaders)

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

	fmt.Println(response.StatusCode)
	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := httpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

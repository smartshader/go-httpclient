package main

import (
	"fmt"
	"github.com/smartshader/go-httpclient/gohttp"
	"io/ioutil"
)

func main() {
	client := gohttp.New()
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

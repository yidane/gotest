package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://36.250.74.88:8103")
	}
	transp := http.Transport{Proxy: proxy}
	client := http.Client{Transport: &transp}
	resp, err := client.Get("http://www.google.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}

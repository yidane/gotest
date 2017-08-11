package main

import (
	"fmt"
	"net/http"
	"errors"
)

func main() {
	url := "http://dwz.cn/5CFTHF"
	location, err := GetRedirectUrl(url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(location)
}

func GetRedirectUrl(url string) (string, error) {
	var client = http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == 302 {
		if location, ok := res.Header["Location"]; ok {
			return location[0], nil
		}
	}

	return "", errors.New("The Response StatusCode is not 302")
}

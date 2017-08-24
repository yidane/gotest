package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetRedirectURL(url string) (string, error) {
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

func GetLongUrl(url string) (*string, error) {
	var client = http.Client{}
	body := strings.NewReader(url)
	req, err := http.NewRequest("POST", "http://dwz.cn/query.php", body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type dwzRes struct {
		Status  int
		LongURL string
	}

	var r dwzRes
	err = json.Unmarshal(bs, &r)
	if err != nil {
		return nil, err
	}

	return &(r.LongURL), nil
}

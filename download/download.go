package download

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func Download() (bool, error) {
	resp, err := http.Get("http://127.0.0.1:8022/")
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	flag := string(body) == "hello"
	if flag {
		return true, nil
	}

	return false, errors.New("Response Body is:" + string(body))
}

package main

import (
	"fmt"
	"net/http"
	"errors"
)

func main() {

	for i := 100000000; i < 100000001; i++ {
		url := "http://dwz.cn/" + Generate(i)
		location, err := GetRedirectUrl(url)
		if err != nil {
			fmt.Println(url)
			fmt.Println(err)
		}

		fmt.Println(location)
	}

	//url := "http://dwz.cn/5CFTHF"
	//location, err := GetRedirectUrl(url)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(location)
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

func Generate(num int) (tiny string) {
	fmt.Println(num)
	alpha := merge(getRange(48, 57), getRange(65, 90))
	alpha = merge(alpha, getRange(97, 122))
	if num < 62 {
		tiny = string(alpha[num])
		return tiny
	} else {
		var runes []rune
		runes = append(runes, alpha[num%62])
		num = num / 62
		for num >= 1 {
			if num < 62 {
				runes = append(runes, alpha[num-1])
			} else {
				runes = append(runes, alpha[num%62])
			}
			num = num / 62

		}
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		tiny = string(runes)
		return tiny
	}
	return tiny
}

func getRange(start, end rune) (ran []rune) {
	for i := start; i <= end; i++ {
		ran = append(ran, i)
	}
	return ran
}

func merge(a, b []rune) []rune {
	c := make([]rune, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}

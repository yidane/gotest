package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

//URLInfo store 302 info
type URLInfo struct {
	ID       int
	Code     string
	Location string
}

var (
	codeChan         = make(chan *URLInfo, 1000)
	locationInfoChan = make(chan *URLInfo, 100)
	total            = 10
)

func init() {
	connectionAndCreateIndex(true)
}

func main() {
	cpuNum := runtime.NumCPU() * 5
	runtime.GOMAXPROCS(cpuNum)

	wait := sync.WaitGroup{}
	wait.Add(cpuNum + 2)
	go func() {
		log.Println("begin generate code")
		for i := 0; i < total; i++ {
			newCode := generate(i)
			codeChan <- &URLInfo{ID: i, Code: newCode, Location: ""}
		}
		wait.Done()
	}()

	getRedirect := func(i int) {
		log.Println("goruntine ", i, " start")
		t := 0
		for t < total {
			urlInfo := <-codeChan
			url := "http://dwz.cn/" + urlInfo.Code
			location, err := getRedirectURL(url)
			if err == nil {
				urlInfo.Location = location
			} else {
				urlInfo = nil
			}
			locationInfoChan <- urlInfo
			t++
		}
		wait.Done()
	}

	for i := 0; i < cpuNum; i++ {
		go getRedirect(i)
	}

	log.Println("start goruntine save to redis")
	time.Sleep(time.Second)

	go func() {
		t := 0
		for t < total {
			urlInfo := <-locationInfoChan
			if urlInfo != nil {
				fmt.Println(urlInfo)
			}
			t++
		}

		wait.Done()
	}()

	wait.Wait()
	fmt.Println("FINISHED")
}

func getRedirectURL(url string) (string, error) {
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
		if location, ok := res.Header["Location"]; ok && len(location) > 0 {
			return location[0], nil
		}
	}

	log.Println(res.StatusCode)
	return "", errors.New("The Response StatusCode is not 302")
}

func generate(num int) (tiny string) {
	alpha := merge(getRange(48, 57), getRange(65, 90))
	alpha = merge(alpha, getRange(97, 122))
	if num < 62 {
		tiny = string(alpha[num])
		return tiny
	}
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

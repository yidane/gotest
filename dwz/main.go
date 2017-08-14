package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
)

//LocationInfo store 302 info
type LocationInfo struct {
	Key      string
	Location string
}

var (
	codeChan         = make(chan string, 1000)
	locationInfoChan = make(chan LocationInfo, 100)
)

func main() {

	con, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer con.Close()
	log.Println("connect redis success")

	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)

	go func() {
		log.Println("begin generate code")
		for i := 0; i < 100000001; i++ {
			newCode := generate(i)
			codeChan <- newCode
		}
	}()

	getRedirect := func(i int) {
		log.Println("goruntine ", i, " start")
		for {
			select {
			case code := <-codeChan:
				url := "http://dwz.cn/" + code
				location, err := getRedirectURL(url)
				if err != nil {
					fmt.Println(url)
					fmt.Println(err)
				}

				locationInfoChan <- LocationInfo{Location: location, Key: code}
				break
			default:
				break
			}
		}
	}

	for i := 0; i < cpuNum; i++ {
		go getRedirect(i)
	}

	log.Println("start goruntine save to redis")
	time.Sleep(time.Second)
	for {
		select {
		case locationInfo := <-locationInfoChan:
			con.Do("SET", locationInfo.Key, locationInfo.Location)
			break
		default:
			break
		}
	}
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
		if location, ok := res.Header["Location"]; ok {
			return location[0], nil
		}
	}

	return "", errors.New("The Response StatusCode is not 302")
}

func generate(num int) (tiny string) {
	fmt.Println(num)
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

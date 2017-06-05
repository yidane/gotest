package benchmark

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type BResult struct {
	Flag     bool
	Duration int64
}

func Test(url string, b int) {
	chn := make(chan BResult, b)
	for i := 1000; i < 1000+b; i++ {
		go get(url+strconv.Itoa(i), chn)
	}

	error := []BResult{}
	success := []BResult{}
	for i := 0; i < b; i++ {
		r := <-chn
		if r.Flag {
			success = append(success, r)
		} else {
			error = append(error, r)
		}
	}

	fmt.Println("成功数量：", len(success))
	fmt.Println("失败数量：", len(error))
}

func get(url string, chn chan BResult) {
	now := time.Now()
	_, err := http.Get(url)
	if err != nil {
		chn <- BResult{Flag: false, Duration: 0}
	}

	chn <- BResult{Flag: true, Duration: time.Now().Sub(now).Nanoseconds()}
}

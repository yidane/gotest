package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"testing"
)

func Benchmark_HttpWeb(b *testing.B) {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	chn := make(chan int)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			res, err := http.Get("http://localhost:8888/Get/" + strconv.Itoa(i))
			if err != nil {
				fmt.Println(err)
			}

			defer res.Body.Close()
			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}

			chn <- i
		}(i)
	}

	for i := 0; i < b.N; i++ {
		<-chn
	}
}

package download

import (
	"fmt"
	"net/http"
	"runtime"
	"testing"
)

// func Test_Download(t *testing.T) {
// 	result, err := Download()
// 	if !result {
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		t.Error("unkonw error")
// 	}
// }

// func Benchmark_Download(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Download()
// 	}
// }

// func Test_Visit(t *testing.T) {
// 	for i := 0; i < 10000; i++ {
// 		visit()
// 	}
// }

// func visit() {
// 	_, err := http.Get("http://www.kcaogen.top/blogInfo?blogId=9")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func Test_Chanel(t *testing.T) {
	fmt.Println("start")
	chn := make(chan Result)
	cpuSum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuSum)
	for i := 0; i < cpuSum; i++ {
		go ChanelDownlowd("http://www.kcaogen.top/blogInfo?blogId=9", i, chn)
		//go ChanelDownlowd("http://localhost/TestWeb/api/tick", i, chn)
	}

	for i := 0; i < cpuSum*10000; i++ {
		fmt.Println(<-chn)
	}
}

func ChanelDownlowd(url string, id int, chn chan Result) {
	for i := 0; i < 10000; i++ {
		http.Get(url)
		result := Result{ID: id, Count: i}
		chn <- result
	}
}

type Result struct {
	ID    int
	Count int
}

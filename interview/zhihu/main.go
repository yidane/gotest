package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i1: ", i, "	", &i)
			wg.Done()
		}()

		runtime.Gosched()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i2: ", i, "	", &i)
			wg.Done()
		}(i)
		runtime.Gosched()
	}
	wg.Wait()

	time.Sleep(time.Second)
}

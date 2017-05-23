package main

import (
	"fmt"
	"runtime"
)

//会出现输出争抢情况 concurrent2:01201234567893456789
func concurrent2() {
	runtime.GOMAXPROCS(2)
	done := make(chan bool)
	go loop2(done)
	go loop2(done)

	<-done
	<-done

	fmt.Println()
}

func loop2(done chan bool) {
	for i := 0; i < max; i++ {
		fmt.Print(i)
	}

	done <- true
}

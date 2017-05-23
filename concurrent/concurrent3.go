package main

import (
	"fmt"
	"runtime"
)

//concurrent3:00112233445566778899
func concurrent3() {
	done := make(chan bool)
	go loop3(done)
	go loop3(done)

	<-done
	<-done
}

func loop3(done chan bool) {
	for i := 0; i < max; i++ {
		fmt.Print(i)
		runtime.Gosched() //让出CPU执行时间
	}

	done <- true
}

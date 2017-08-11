package channel

import (
	"fmt"
	"time"
)

func mainFunc() {
	requestChan := make(chan chan string)
	go requestFunc(requestChan)
	go responseFunc(requestChan)
	time.Sleep(time.Second)
}

func requestFunc(c chan chan string) {
	requestChan := make(chan string)
	c <- requestChan
	request := <-requestChan
	fmt.Println(request)
}

func responseFunc(c chan chan string) {
	requestChan := <-c
	requestChan <- "hello"
}

func timeChan() {
	//go t()
	chl1 := make(chan int, 1)
	sign := make(chan byte, 1)

	go func() {
		for i := 0; i < 10; i++ {
			chl1 <- 1
		}
	}()

	go func() {
		var e int
		ok := true

		var timer *time.Timer

		for {
			select {
			case e = <-chl1:
				fmt.Printf("ch1->%d\n", e)
				fmt.Println(time.Now())
			case t := <-func() <-chan time.Time {
				if timer == nil {
					timer = time.NewTimer(2 * time.Second)
				} else {
					timer.Reset(2 * time.Second)
				}
				return timer.C
			}():
				fmt.Println("Timeout.", t)
				ok = false
				break
			}

			if !ok {
				sign <- 0
				break
			}
		}
	}()

	<-sign
}

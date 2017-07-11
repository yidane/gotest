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

package main

import (
	"fmt"
	"time"
)

var t int64

func producer(ch chan int64) {
	for {
		t++
		ch <- t
	}
}

func consumer(ch chan int64) {
	for {
		fmt.Println(<-ch)
	}
}

func producerAndConsumer() {
	ch := make(chan int64)
	go producer(ch)
	go consumer(ch)
	time.Sleep(time.Minute)
}

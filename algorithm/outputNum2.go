package main

import "time"

func output2() {
	cha := make(chan struct{})
	chb := make(chan struct{})
	chc := make(chan struct{})
	chd := make(chan struct{})
	go test(1, cha, chb)
	go test(2, chb, chc)
	go test(3, chc, chd)
	go test(4, chd, cha)
	cha <- struct{}{}
	time.Sleep(time.Second * 2) //阻塞
}

func test(i int, ch1, ch2 chan struct{}) {
	for range ch1 {
		println(i)
		ch2 <- struct{}{}
	}
}

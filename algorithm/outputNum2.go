package main

import "time"
import "fmt"

var arr []int = make([]int, 0)

func output2() {
	cha := make(chan []int)
	chb := make(chan []int)
	chc := make(chan []int)
	chd := make(chan []int)
	go test(1, cha, chb)
	go test(2, chb, chc)
	go test(3, chc, chd)
	go test(4, chd, cha)
	cha <- []int{}
	time.Sleep(time.Second / 100) //阻塞

	fmt.Println(arr)
}

func test(i int, ch1, ch2 chan []int) {
	for range ch1 {
		arr = append(arr, i)
		ch2 <- []int{}
	}
}

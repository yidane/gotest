package main

import (
	"bytes"
	"log"
	"time"
)

func output1() {
	ch1ss := makeChs(4)
	ch2ss := makeChs(4)
	ch3ss := makeChs(4)
	ch4ss := makeChs(4)

	go gen(1, ch1ss)
	go gen(2, ch2ss)
	go gen(3, ch3ss)
	go gen(4, ch4ss)

	chs1 := []chan int{ch1ss[0], ch2ss[0], ch3ss[0], ch4ss[0]}
	buf1 := new(bytes.Buffer)
	go writeFile(chs1, buf1)

	chs2 := []chan int{ch2ss[1], ch3ss[1], ch4ss[1], ch1ss[1]}
	buf2 := new(bytes.Buffer)
	go writeFile(chs2, buf2)

	chs3 := []chan int{ch3ss[2], ch4ss[2], ch1ss[2], ch2ss[2]}
	buf3 := new(bytes.Buffer)
	go writeFile(chs3, buf3)

	chs4 := []chan int{ch4ss[3], ch1ss[3], ch2ss[3], ch3ss[3]}
	buf4 := new(bytes.Buffer)
	go writeFile(chs4, buf4)

	time.Sleep(1 * time.Second)
	printData("file1: ", buf1)
	printData("file2: ", buf2)
	printData("file3: ", buf3)
	printData("file4: ", buf4)
}

func makeChs(len int) (rets []chan int) {
	rets = make([]chan int, len)
	for i := 0; i < len; i++ {
		rets[i] = make(chan int)
	}
	return
}

func gen(seed int, rets []chan int) {
	for {
		for _, ch := range rets {
			ch <- seed
			time.Sleep(time.Millisecond)
		}
	}
}

func writeFile(chs []chan int, buf *bytes.Buffer) {
	for {
		for _, ch := range chs {
			data := <-ch
			buf.WriteRune(rune(data))
		}
	}
}

func printData(prefix string, buf *bytes.Buffer) {
	log.Println(prefix, buf.Bytes())
}

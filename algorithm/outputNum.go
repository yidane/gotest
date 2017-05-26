package main

import "fmt"
import "bytes"
import "time"

// 有四个线程1、2、3、4。线程1的功能就是输出1，线程2的功能就是输出2，以此类推.........
// 现在有四个文件ABCD。初始都为空。现要让四个文件呈如下格式：
// A：1 2 3 4 1 2....
// B：2 3 4 1 2 3....
// C：3 4 1 2 3 4....
// D：4 1 2 3 4 1....
func output() {
	A := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}
	B := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}
	C := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}
	D := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}

	go writeData(1, A)
	go writeData(2, B)
	go writeData(3, C)
	go writeData(4, D)

	buf1 := new(bytes.Buffer)
	go boxData(buf1, A[0], B[0], C[0], D[0])
	buf2 := new(bytes.Buffer)
	go boxData(buf2, A[1], B[1], C[1], D[1])
	buf3 := new(bytes.Buffer)
	go boxData(buf3, A[2], B[2], C[2], D[2])
	buf4 := new(bytes.Buffer)
	go boxData(buf4, A[3], B[3], C[3], D[3])

	time.Sleep(time.Second)
	fmt.Println(buf1)
	fmt.Println(buf2)
	fmt.Println(buf3)
	fmt.Println(buf4)
}

func writeData(i int, chanArr []chan int) {
	for {
		for _, chn := range chanArr {
			chn <- i
		}
	}
}

func boxData(buf *bytes.Buffer, chanArr ...chan int) {
	for {
		for _, chn := range chanArr {
			data := <-chn
			buf.WriteRune(rune(data))
		}
	}
}

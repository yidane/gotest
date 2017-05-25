package main

import (
	"fmt"
)

// 有四个线程1、2、3、4。线程1的功能就是输出1，线程2的功能就是输出2，以此类推.........
// 现在有四个文件ABCD。初始都为空。现要让四个文件呈如下格式：
// A：1 2 3 4 1 2....
// B：2 3 4 1 2 3....
// C：3 4 1 2 3 4....
// D：4 1 2 3 4 1....
func output(max int) {
	A := make(chan int)
	B := make(chan int)
	C := make(chan int)
	D := make(chan int)

	arrA := []int{}
	arrB := []int{}
	arrC := []int{}
	arrD := []int{}

	go func() {
		times := 0
		for {
			times++
		}
	}()

	go func() {
		
	}()

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
}

func add1(c chan (int)) {
	c <- 1
}
func add2(c chan (int)) {
	c <- 2
}
func add3(c chan (int)) {
	c <- 3
}
func add4(c chan (int)) {
	c <- 4
}

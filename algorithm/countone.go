// 有一个整数n,写一个函数f(n),返回0到n之间出现的"1"的个数。
// 比如f(13)=6,现在f(1)=1,问下一个最大的f(n)=n的n是什么？

package main

import "fmt"
import "time"

var cache map[int]int = make(map[int]int)
var totalCache map[int]int = make(map[int]int)

//Fn 返回0到max之间出现的“1”的个数
func Fn(max int) int {
	if max <= 0 {
		return 0
	}
	if max == 1 {
		totalCache[1] = 1
		return 1
	}
	var result = CountOne(max)
	cache[max] = CountOne(max)
	result += totalCache[max-1]

	totalCache[max] = result
	return result
}

func CountOne(num int) int {
	if _, ok := cache[num]; ok {
		return cache[num]
	}

	if num < 1 {
		return 0
	}
	if num == 1 || num == 10 {
		return 1
	}
	if num < 10 && num > 1 {
		return 0
	}
	var result int
	m := num % 10
	if m == 1 {
		result++
	}
	result = result + CountOne(num/10)

	return result
}

//CountOneExec 输出200000以内所有F(n)=n的数
func CountOneExec() {
	beginTime := time.Now()
	fmt.Println("begin")
	for i := 0; i < 200000; i++ {
		if i == Fn(i) {
			fmt.Println(i)
		}
	}
	fmt.Println("time:", time.Now().Sub(beginTime).String())
	fmt.Println("finish")
}

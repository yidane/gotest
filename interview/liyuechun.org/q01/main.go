package main

import (
"fmt"
)

func main() {
	defer_call3()
	defer_call1()
	defer_call2()
	defer_call()
}

func defer_call() {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

func defer_call1() {

	defer func() {
		fmt.Println("打印前1")
	}()

	defer func() {
		fmt.Println("打印中1")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常

		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印后1")
	}()

	panic("触发异常1")
}

func defer_call2() {

	defer func() {
		fmt.Println("打印前2")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中2")
	}()

	defer func() {

		fmt.Println("打印后2")
	}()

	panic("触发异常2")
}

func defer_call3() {

	defer func() {
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印前3")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中3")
	}()

	defer func() {
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印后3")
	}()

	panic("触发异常3")
}
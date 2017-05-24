package main

import (
	"fmt"
	"time"
)

//Output 不会输出任何值 加上time.Sleep(time.Second)后，会输出值，但是输出值不一定都是10
func Output(times uint) {
	var i uint
	for i = 0; i < times; i++ {
		go func(i uint) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}

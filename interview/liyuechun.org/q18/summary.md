####18. 请说出下面代码哪里写错了。
~~~
func main() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for {
			a := <-abc
			fmt.Println("a: ", a)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}
~~~

go中的for循环是死循环，应该设置出口。正确代码如下：

~~~
package main

import (
	"fmt"
	"time"
)

func main() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for {
			a,ok := <-abc
			if !ok {
				fmt.Println("结束！")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}
~~~
运行结果为：

~~~
close
a:  0
a:  1
a:  2
a:  3
a:  4
a:  5
a:  6
a:  7
a:  8
a:  9
结束！
~~~
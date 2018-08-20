####15. 请找出下面代码的问题所在。
~~~
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
~~~
运行结果如下：
~~~
panic: send on closed channel
ok

goroutine 5 [running]:
main.main.func1(0xc420098000)
~~~
解析：出现上面错误的原因是因为提前关闭通道所致。

正确代码如下：
~~~
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个缓冲通道
	ch := make(chan int, 1000)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			a, ok := <-ch
			
			if !ok {
				fmt.Println("close")
				close(ch)
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	fmt.Println("ok")
	time.Sleep(time.Second)
}
~~~
运行结果如下：
~~~
ok
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
~~~
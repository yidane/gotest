####1. 写出下面代码输出内容
~~~
package main

import (
	"fmt"
)

func main() {
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
~~~
在这个案例中，触发异常这几个字打印的顺序其实是不确定的。defer, panic, recover一般都会配套使用来捕捉异常。
<br/>先看下面的案例：

#####案例一
~~~
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常

		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")
}
~~~
输出内容为：
~~~
触发异常
打印后
打印中
打印前

Process finished with exit code 0
~~~
案例二
~~~
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	defer func() {

		fmt.Println("打印后")
	}()

	panic("触发异常")
}
~~~
输出内容为：
~~~
打印后
触发异常
打印中
打印前

Process finished with exit code 0
~~~
案例三
~~~
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {

	defer func() {
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印前")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	defer func() {
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")
}
~~~
输出内容为：
~~~
触发异常
打印后
打印中
打印前

Process finished with exit code 0
~~~

####总结：

1、defer函数属延迟执行，延迟到调用者函数执行 return 命令前被执行。<br>
2、多个defer之间按LIFO先进后出顺序执行。<br>
3、Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
如果同时有多个defer，那么异常会被最近的recover()捕获并正常处理。
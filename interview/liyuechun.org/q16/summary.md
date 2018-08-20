####16. 请说明下面代码书写是否正确。
~~~
var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v(v+delta)) {
			break
		}
	}
}
~~~
atomic.CompareAndSwapInt32里面一共三个参数，上面的书写错误，正确的书写是：
~~~
atomic.CompareAndSwapInt32(&value, v,v+delta)
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
~~~
第一个参数的值应该是指向被操作值的指针值。该值的类型即为*int32。<br>
后两个参数的类型都是int32类型。它们的值应该分别代表被操作值的旧值和新值
CompareAndSwapInt32·函数在被调用之后会先判断参数addr指向的被操作值与参数old`的值是否相等。
仅当此判断得到肯定的结果之后，该函数才会用参数new代表的新值替换掉原先的旧值。否则，后面的替换操作就会被忽略。
完整代码如下：
~~~
package main

import (
	"sync/atomic"
	"fmt"
)

var value int32

func SetValue(delta int32) {
	for {
		v := value
		// 比较并交换

		if atomic.CompareAndSwapInt32(&value, v,v+delta) {
			fmt.Println(value)
			break
		}
	}
}

func main()  {
	SetValue(100)
}
~~~
运行结果为100.
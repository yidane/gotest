####14. 下面的代码是有问题的，请说明原因。
~~~
package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	p.String()
}
~~~

运行结果如下：
~~~
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow

runtime stack:
runtime.throw(0x10c122b, 0xe)
~~~
如下所示，上面的代码出现了栈溢出，原因是因为%v格式化字符串是本身会调用String()方法，上面的栈溢出是因为无限递归所致。
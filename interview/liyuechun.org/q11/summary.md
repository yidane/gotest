####11. 下面代码能运行吗？为什么
~~~
type Param map[string]interface{}

type Show struct {
	Param
}

func main1() {
	s := new(Show)
	s.Param["RMB"] = 10000
}
~~~
运行结果：

~~~
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
~~~
如上所示，运行过程中会发生异常，原因是因为字典Param的默认值为nil，当给字典nil增加键值对是就会发生运行时错误panic: assignment to entry in nil map。

正确的修改方案如下：
~~~
package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {

	// 创建Show结构体对象
	s := new(Show)
	// 为字典Param赋初始值
	s.Param = Param{}
	// 修改键值对
	s.Param["RMB"] = 10000
	fmt.Println(s)
}
~~~
运行结果如下：
~~~
&{map[RMB:10000]}

Process finished with exit code 0
~~~
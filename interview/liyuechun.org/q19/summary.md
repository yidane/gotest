####19. 请说出下面代码，执行时为什么会报错
~~~
type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"liyuechun"}}
	m["people"].name = "wuyanzu"
}
~~~
答案：报错的原因是因为不能修改字典中value为结构体的属性值。

代码作如下修改方可运行：
~~~
package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"liyuechun"}}
	fmt.Println(m)
	fmt.Println(m["people"])

	// 不能修改字典中结构体属性的值
	//m["people"].name = "wuyanzu"
	
	var s Student = m["people"] //深拷贝
	s.name = "xietingfeng"
	fmt.Println(m)
	fmt.Println(s)
}
~~~
运行结果如下：
~~~
map[people:{liyuechun}]
{liyuechun}
map[people:{liyuechun}]
{wuyanzu}
~~~
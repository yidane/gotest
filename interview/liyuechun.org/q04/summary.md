~~~
package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
~~~
运行结果如下：
~~~
showA
showB

Process finished with exit code 0
~~~
Go中没有继承,上面这种写法叫组合。

上面的t.ShowA()等价于t.People.ShowA()，将上面的代码修改如下：
~~~
func main() {
	t := Teacher{}
	t.ShowA()
	fmt.Println("---------------")
	t.People.ShowA()
}
~~~
运行结果为：
~~~
showA
showB
---------------
showA
showB

Process finished with exit code 0
~~~
####10. 以下代码能编译过去吗？为什么？
~~~
package main
import (
	"fmt"
)
type People interface {
	Speak(string) string
}
type Stduent struct{}
func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
~~~
结论：编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People 类型。

两种正确修改方法：

方法一
~~~
package main
import (
	"fmt"
)
type People interface {
    Speak(string) string
}

type Stduent struct{}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	var peo People = Stduent{}
	think := "hi"
	fmt.Println(peo.Speak(think))
}
~~~
方法二
~~~
package main
import (
	"fmt"
)
type People interface {
    Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
~~~
总结：指针类型的结构体对象可以同时调用结构体值类型和指针类型对应的方法。而值类型的结构体对象只能调用值类型对应的接口方法。


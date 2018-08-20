####2. 以下代码有什么问题，说明原因
~~~
package main
import (
	"fmt"
)
type student struct {
	Name string
	Age  int
}
func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	return m
}
func main() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
~~~
运行结果：
~~~
key=zhou,value=&{wang 22} 
key=li,value=&{wang 22} 
key=wang,value=&{wang 22} 

Process finished with exit code 0
~~~
修改一下代码：

将下面的代码：
~~~
for _, stu := range stus {
    m[stu.Name] = &stu
}
~~~
修改为：
~~~
for _, stu := range stus {
	fmt.Printf("%v\t%p\n",stu,&stu)
	m[stu.Name] = &stu
}
~~~
运行结果为：
~~~
{shen 24}	0xc4200a4020
{li 23}	0xc4200a4020
{wang 22}	0xc4200a4020
key=shen,value=&{wang 22} 
key=li,value=&{wang 22} 
key=wang,value=&{wang 22} 

Process finished with exit code 0
~~~
通过上面的案例，我们不难发现stu变量的地址始终保持不变，每次遍历仅进行struct值拷贝，故m[stu.Name]=&stu实际上一直指向同一个地址，最终该地址的值为遍历的最后一个struct的值拷贝。

形同如下代码：
~~~
var stu student 
for _, stu = range stus {
	m[stu.Name] = &stu
} 
~~~
修正方案，取数组中原始值的地址：
~~~
for i, _ := range stus {
	stu:=stus[i]
	m[stu.Name] = &stu
}
~~~
重新运行，效果如下：
~~~
{shen 24}	0xc42000a060
{li 23}	0xc42000a0a0
{wang 22}	0xc42000a0e0
key=shen,value=&{shen 24} 
key=li,value=&{li 23} 
key=wang,value=&{wang 22} 

Process finished with exit code 0
~~~

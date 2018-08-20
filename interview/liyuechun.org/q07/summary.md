####7. 请写出以下输入内容
~~~
package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Printf("%p\n", s)
	s = append(s, 1, 2, 3)
	fmt.Printf("%p\n", s) //new pointer
	fmt.Println(s)
}
~~~
运行结果:

~~~
0xc4200180c0
0xc42001c0a0
[0 0 0 0 0 1 2 3]

Process finished with exit code 0
~~~
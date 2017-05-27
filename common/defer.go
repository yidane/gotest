package main

import "fmt"

//http://studygolang.com/articles/4809
//多个defer的执行顺序为“后进先出”；
//defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i) // 打印结果为 c defer: 2
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i) // 打印结果为 c defer: 1
	}()
	return &i
}

func deferFunc() {
	fmt.Println("A:")
	fmt.Println("return:", a())
	fmt.Println("B:")
	fmt.Println("return:", b())
	fmt.Println("C:")
	fmt.Println("c return:", *(c()))
}

/*
A:
defer1: 1
defer2: 2
return: 0
B:
defer1: 1
defer2: 2
return: 2
C:
c defer1: 1
c defer2: 2
c return: 2

先来假设出结论，帮助大家理解原因：
	多个defer的执行顺序为“后进先出”；
	defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。
如何解释两种结果的不同：
	上面两段代码的返回结果之所以不同，其实从上面第2条结论很好理解。
	a()int 函数的返回值没有被提前声名，其值来自于其他变量的赋值，而defer中修改的也是其他变量，而非返回值本身，因此函数退出时返回值并没有被改变。
	b()(i int) 函数的返回值被提前声名，也就意味着defer中是可以调用到真实返回值的，因此defer在return赋值返回值 i 之后，再一次地修改了 i 的值，
		最终函数退出后的返回值才会是defer修改过的值。
	c()*int 的返回值没有被提前声明，但是由于 c()*int 的返回值是指针变量，那么在return将变量 i 的地址赋给返回值后，defer再次修改了 i 在内存中的实际值，
		因此函数退出时返回值虽然依旧是原来的指针地址，但是其指向的内存实际值已经被成功修改了。
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	makeAppend()
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")

	/*
		打印后
		打印中
		打印前
		panic: 触发异常

		goroutine 1 [running]:
		main.deferCall()
				/home/yidane/Workspace/go/src/github.com/yidane/gotest/interview/csdn/main.go:15 +0xc0
		main.main()
				/home/yidane/Workspace/go/src/github.com/yidane/gotest/interview/csdn/main.go:6 +0x26
		exit status 2
	*/
}

func forRange() {
	type student struct {
		Name string
		Age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Printf("%p\n", &stu)
	}

	fmt.Println(m)
	/*
		这样的写法初学者经常会遇到的，很危险！ 与Java的foreach一样，都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝。 就像想修改切片元素的属性：

		for _, stu := range stus {
		    stu.Age = stu.Age+10
		}
		也是不可行的。 大家可以试试打印出来：

		func pase_student() {
		    m := make(map[string]*student)
		    stus := []student{
		        {Name: "zhou", Age: 24},
		        {Name: "li", Age: 23},
		        {Name: "wang", Age: 22},
		    }
		    // 错误写法
		    for _, stu := range stus {
		        m[stu.Name] = &stu
		    }

		    for k,v:=range m{
		        println(k,"=>",v.Name)
		    }

		    // 正确
		    for i:=0;i<len(stus);i++  {
		        m[stus[i].Name] = &stus[i]
		    }
		    for k,v:=range m{
		        println(k,"=>",v.Name)
		    }
		}
	*/
}

func goOutput() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	/*
		谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。 但是A:均为输出10，B:从0~9输出(顺序不定)。 第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。 故go func执行时，i的值始终是10。

		第二个go func中i是函数参数，与外部for中的i完全是两个变量。 尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
	*/
}

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

func oppFunc() {
	t := Teacher{}
	t.ShowA()
	t.ShowB()

	/*
		这是Golang的组合模式，可以实现OOP的继承。 被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），但它们的方法(ShowA())调用时接受者并没有发生变化。 此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
	*/
}

func chanFunc() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}

	/*
		select会随机选择一个可用通用做收发操作。 所以代码是有肯触发异常，也有可能不会。 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：
		select 中只要有一个case能return，则立刻执行。
		当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
		如果没有一个case能return则可以执行”default”块。
	*/
}

func deferFunc() {
	calc := func(index string, a, b int) int {
		ret := a + b
		fmt.Println(index, a, b, ret)
		return ret
	}

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	/*
		这道题类似第1题 需要注意到defer执行顺序和值传递 index:1肯定是最后执行的，但是index:1的第三个参数是一个函数，所以最先被调用calc("10",1,2)==>10,1,2,3 执行index:2时,与之前一样，需要先调用calc("20",0,2)==>20,0,2,2 执行到b=1时候开始调用，index:2==>calc("2",0,2)==>2,0,2,2 最后执行index:1==>calc("1",1,3)==>1,1,3,4

		10 1 2 3
		20 0 2 2
		2 0 2 2
		1 1 3 4
	*/
}

func makeAppend() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	/*
		make初始化是由默认值的哦，此处默认值为0
		[0 0 0 0 0 1 2 3]
	*/
}

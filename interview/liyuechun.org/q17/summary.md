####17. 下面的程序运行后为什么会爆异常。
~~~
package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 100)
}

func main() {
	p := new(Project)
	p.Main()
}
~~~
运行结果如下：
~~~
panic: interface conversion: interface {} is string, not int

goroutine 17 [running]:
main.(*Project).exec(0x1157c08, 0xc420068060)
~~~
出现异常的原因是因为写入到管道的数据类型为string,而m := msg.(int)这句代码里面却使用了int，修改方法，将int修改为string即可。

完整正确代码如下：
~~~
package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		m := msg.(string)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 100)
}

func main() {
	p := new(Project)
	p.Main()
}
~~~
运行结果如下：
~~~
msg:  1
msg:  1
msg:  1
.
.
.
msg:  1
msg:  1
msg:  1
msg:  1
msg:  1
~~~
####9. 下面的迭代会有什么问题？
~~~
package main

import "fmt"
import "sync"
import "time"

type ThreadSafeSet struct {
	sync.RWMutex
	s []int
}

func (set *ThreadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
			fmt.Print("get:", elem, ",")
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {
	//read()
	unRead()
}
func read() {
	set := ThreadSafeSet{}
	set.s = make([]int, 100)
	ch := set.Iter()
	closed := false
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Print("read:", v, ",")
			} else {
				closed = true
			}
		}
		if closed {
			fmt.Print("closed")
			break
		}
	}
	fmt.Print("Done")
}

func unRead() {
	set := ThreadSafeSet{}
	set.s = make([]int, 100)
	ch := set.Iter()
	_ = ch
	time.Sleep(5 * time.Second)
	fmt.Print("Done")
}
~~~

结论：内部迭代出现阻塞。默认初始化时无缓冲区，需要等待接收者读取后才能继续写入。

chan在使用make初始化时可附带一个可选参数来设置缓冲区。默认无缓冲，题目中便初始化的是无缓冲区的chan，这样只有写入的元素直到被读取后才能继续写入，不然就一直阻塞。

设置缓冲区大小后，写入数据时可连续写入到缓冲区中，直到缓冲区被占满。从chan中接收一次便可从缓冲区中释放一次。可以理解为chan是可以设置吞吐量的处理池。
~~~
ch := make(chan interface{}) 和ch := make(chan interface{},1)是不一样的 
无缓冲的 不仅仅是只能向 ch 通道放 一个值 而是一直要有人接收，
那么ch <- elem才会继续下去，要不然就一直阻塞着，也就是说有接收者才去放，
没有接收者就阻塞。

而缓冲为1则即使没有接收者也不会阻塞，因为缓冲大小是1只有当 
放第二个值的时候 第一个还没被人拿走，这时候才会阻塞
~~~
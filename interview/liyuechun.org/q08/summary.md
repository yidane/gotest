####8. 下面的代码有什么问题
~~~
package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	count := 1000
	gw := sync.WaitGroup{}
	gw.Add(count * 3)
	u := UserAges{ages: map[string]int{}}
	add := func(i int) {
		u.Add(fmt.Sprintf("user_%d", i), i)
		gw.Done()
	}
	
	for i := 0; i < count; i++ {
		go add(i)
		go add(i)
	}
	
	for i := 0; i < count; i++ {
		go func(i int) {
			defer gw.Done()
			u.Get(fmt.Sprintf("user_%d", i))
		}(i)
	}
	gw.Wait()
	fmt.Println("Done")
}
~~~
输出结果：
~~~
fatal error: concurrent map read and map write

goroutine 2022 [running]:
runtime.throw(0x10c5472, 0x21)
~~~
结论： 在执行 Get 方法时可能被panic。

虽然有使用sync.Mutex做写锁，但是map是并发读写不安全的。map属于引用类型，并发读写时多个协程见是通过指针访问同一个地址，即访问共享变量，此时同时读写资源存在竞争关系。所以会报错误信息:fatal error: concurrent map read and map write。

如果第一次没复现panic问题，可以再次运行，复现该问题。那么如何改善呢? 在Go1.9新版本中将提供并发安全的map。首先需要了解两种锁的不同：

sync.Mutex互斥锁
sync.RWMutex读写锁，基于互斥锁的实现，可以加多个读锁或者一个写锁。
RWMutex相关方法：
~~~
type RWMutex
    func (rw *RWMutex) Lock() 
    func (rw *RWMutex) RLock()
    func (rw *RWMutex) RLocker() Locker
    func (rw *RWMutex) RUnlock()
    func (rw *RWMutex) Unlock()
~~~
代码改进如下：
~~~
package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}

	return -1
}

func main() {
	count := 10000
	gw := sync.WaitGroup{}
	gw.Add(count * 3)
	u := UserAges{ages: map[string]int{}}
	add := func(i int) {
		u.Add(fmt.Sprintf("user_%d", i), i)
		gw.Done()
	}
	for i := 0; i < count; i++ {
		go add(i)
		go add(i)
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			defer gw.Done()
			u.Get(fmt.Sprintf("user_%d", i))
			fmt.Print(".")
		}(i)
	}
	gw.Wait()
	fmt.Println("Done")
}
~~~
运行结果如下：
~~~
.
.
.
.
Done

Process finished with exit code 0
~~~
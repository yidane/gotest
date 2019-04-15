package syncTest

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test_atomic(t *testing.T) {
	var i int32
	for ; i < 10; i++ {
		go add(i)
	}

	time.Sleep(time.Millisecond * 100)

	fmt.Println(get())
}

func Test_Cond(t *testing.T) {
	lock := sync.RWMutex{}
	var cond = sync.NewCond(&lock)

	cond.Signal()
	cond.Broadcast()
	cond.Signal()

	lock.RLock()
	defer lock.RUnlock()
	go func() {
		lock.Lock()
		defer lock.Unlock()
	}()

	runtime.GC()
	fmt.Println(runtime.GOROOT())
}

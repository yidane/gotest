package sync

import (
	"fmt"
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

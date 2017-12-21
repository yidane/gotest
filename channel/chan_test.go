package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_ChanChan(t *testing.T) {
	mainFunc()
}

func Test_TimeChan(t *testing.T) {
	timeChan()
}

func Test_BufferChan(t *testing.T) {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("GO!")
		v := <-c
		fmt.Println(v)
	}()

	c <- true
	fmt.Println("222")
}

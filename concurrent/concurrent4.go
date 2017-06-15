package main

import (
	"fmt"
	"sync"
)

func func1() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan string)
	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()

	ch <- "yidane"
	wg.Wait()
}

func func2() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
	for {
		data, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(data)
	}
}

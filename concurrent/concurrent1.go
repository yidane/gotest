package main

import "fmt"

//顺序输出 concurrent1:01234567890123456789
func concurrent1() {
	done := make(chan bool)
	go loop1(done)
	go loop1(done)

	<-done
	<-done

	fmt.Println("")
}

func loop1(done chan bool) {
	for i := 0; i < max; i++ {
		fmt.Print(i)
	}

	done <- true
}

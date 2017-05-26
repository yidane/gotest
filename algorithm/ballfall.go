package main

import (
	"fmt"
)

func ballfall() {
	var h float64 = 100
	t := 10
	i := 1
	var total float64 = 0

	for i <= t {
		total += h
		h = h / 2
		i++
	}

	fmt.Println("第十次高度: ", h)
	fmt.Println("总经过高度: ", total)
}

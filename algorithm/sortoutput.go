package main

import (
	"fmt"
)

func sortOutput(a, b, c int) {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}

	fmt.Println(a, b, c)
}

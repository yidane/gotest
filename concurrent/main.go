package main

import (
	"fmt"
)

var max int = 10

func main() {
	fmt.Print("concurrent1:")
	concurrent1()
	fmt.Print("concurrent2:")
	concurrent2()
	fmt.Print("concurrent3:")
	concurrent3()
}

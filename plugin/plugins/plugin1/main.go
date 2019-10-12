package main

import "fmt"

func Say(str string) string {
	return fmt.Sprintf("Hello, my name is: %s, addr is: %s\n", str, Addr)
}

var Addr = "default home address"

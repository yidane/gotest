package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("cmd", "w32tm /resync")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

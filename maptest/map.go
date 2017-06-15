package maptest

import "fmt"

func testRefernceType() {
	slice0 := make([]int, 4, 5)
	fmt.Println(slice0)
	fmt.Printf("%p\n", &slice0)
	slice0 = append(slice0, 1)
	fmt.Printf("%p\n", &slice0)
}

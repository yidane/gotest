package main

import "fmt"
import "reflect"

func testArgument() {
	//切片，引用类型
	a := []int{1, 2, 3, 4}
	b := a
	a[1] = 5
	fmt.Println("a	", a, "	b	", b, "	", b[1], "a:", reflect.TypeOf(a), "b:", reflect.TypeOf(b))

	//不定长度的数组，只初始化初始值内容。值类型
	a1 := [...]int{1, 2, 3, 4}
	b1 := a1
	a1[1] = 5
	fmt.Println("a1	", a1, "	b1	", b1, "	", b1[1], "a1:", reflect.TypeOf(a1), "b1:", reflect.TypeOf(b1))

	//定长数组，若初始化值不够数组长度，用0补充。值类型
	a2 := [4]int{1, 2, 3, 4}
	b2 := a2
	a2[1] = 5
	fmt.Println("a2	", a2, "	b2	", b2, "	", b2[1], "a2:", reflect.TypeOf(a2), "b2:", reflect.TypeOf(b2))
}

package main

import "testing"
import "fmt"

func Test_CountOne(t *testing.T) {
	if CountOne(1) == 1 {
		t.Fatal("1:", 1)
	}
	if CountOne(2) == 1 {
		t.Log("2:", 2)
	}
	if CountOne(11) == 3 {
		t.Log("11:", 3)
	}
}

func Test_Fn(t *testing.T) {
	for i := 0; i < 200000; i++ {
		if i == Fn(i) {
			fmt.Println(i)
		}
	}
}

func Test_Sub(t *testing.T) {
	for i := 10; i < 20; i++ {
		fmt.Printf("%d/10=%d\r", i, i/10)
	}

	sub := 10.0
	for i := 10; i < 20; i++ {
		fmt.Printf("%d/10=%f\r", i, float64(i)/sub)
	}

	fmt.Println("格式化输出")
	fmt.Println("http://www.cnblogs.com/golove/p/3284304.html")
	fmt.Println("旗标、宽度、精度、索引")
	fmt.Printf("|%0+- #[1]*.[2]*[3]d|%0+- #[1]*.[2]*[4]d|\n", 8, 4, 32, 64)

	fmt.Println("浮点型精度")
	fmt.Printf("|%f|%8.4f|%8.f|%.4f|%.f|\n", 3.2, 3.2, 3.2, 3.2, 3.2)
	fmt.Printf("|%.3f|%.3g|\n", 12.345678, 12.345678)
	fmt.Printf("|%.2f|\n", 12.345678+12.345678i)

	fmt.Println("字符串精度")
	s := "你好世界！"
	fmt.Printf("|%s|%8.2s|%8.s|%.2s|%.s|\n", s, s, s, s, s)
	fmt.Printf("|%x|%8.2x|%8.x|%.2x|%.x|\n", s, s, s, s, s)

	fmt.Println("带引号字符串")
	s1 := "Hello 世界!"       // CanBackquote
	s2 := "Hello\n世界!"      // !CanBackquote
	fmt.Printf("%q\n", s1)  // 双引号
	fmt.Printf("%#q\n", s1) // 反引号成功
	fmt.Printf("%#q\n", s2) // 反引号失败
	fmt.Printf("%+q\n", s2) // 仅包含 ASCII 字符

	fmt.Println("Unicode 码点")
	fmt.Printf("%U, %#U\n", '好', '好')
	fmt.Printf("%U, %#U\n", '\n', '\n')

	fmt.Println("接口类型将输出其内部包含的值")
	var i interface{} = struct {
		name string
		age  int
	}{"AAA", 20}
	fmt.Printf("%v\n", i)  // 只输出字段值
	fmt.Printf("%+v\n", i) // 同时输出字段名
	fmt.Printf("%#v\n", i) // Go 语法格式

	fmt.Println("输出类型")
	fmt.Printf("%T\n", i)
}

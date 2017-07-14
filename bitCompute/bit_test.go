package bitCompute

import "testing"
import "fmt"

func Test_And(t *testing.T) {
	fmt.Println(3 & 4)
}

func Test_Or(t *testing.T) {
	fmt.Println(3 | 4)
	fmt.Println(3 | -4)
	fmt.Println(3 | 12)
	fmt.Println(10 | 11)
	fmt.Println(8 | 15)
}

func Test_Xor(t *testing.T) {
	fmt.Println(2 ^ 3)
	fmt.Println(4 ^ 5)
	fmt.Println(54 ^ 100)

	//不使用中间变量交换俩整数
	a, b := 4, 5
	fmt.Println(a, b)
	a ^= b
	fmt.Println(a, b)
	b ^= a
	fmt.Println(a, b)
	a ^= b
	fmt.Println(a, b)
}

func Test_Shifting(t *testing.T) {
	fmt.Println(2 << 3)
	fmt.Println(2 >> 3)
}

func Test_Question(t *testing.T) {
	type args struct {
		Title  string
		Result int
	}
	a := 82
	var k uint = 3
	tests := []args{
		{Title: "去掉最后一位", Result: a >> 1},
		{Title: "在最后加一个0", Result: a << 1},
		{Title: "在最后加一个1", Result: a & 1},
		{Title: "把最后一位变成1", Result: a | 1},
		{Title: "最后一位取反", Result: a ^ 1},
		{Title: "把右数第k位变成1", Result: a | (1 << (k - 1))},
		{Title: "把右数第k位变成0", Result: a & (1 << (k - 1))},
		{Title: "右数第k位取反", Result: 0},
		{Title: "取末三位", Result: 0},
		{Title: "取末k位", Result: 0},
		{Title: "取右数第k位", Result: 0},
		{Title: "把末k位变成1", Result: 0},
		{Title: "末k位取反", Result: 0},
		{Title: "把右边连续的1变成0", Result: 0},
		{Title: "把右起第一个0变成1", Result: 0},
		{Title: "把右边连续的0变成1", Result: 0},
		{Title: "取右边连续的1", Result: 0},
		{Title: "去掉右起第一个1的左边", Result: 0},
	}
	fmt.Printf("a:%v\n", a)
	fmt.Printf("k:%v\n", k)
	fmt.Printf("a:%b\n", a)
	for _, arg := range tests {
		t.Run(arg.Title, func(t *testing.T) {
			fmt.Printf("%s:%b\n", arg.Title, arg.Result)
		})
	}
}

//计算一个int型在转换为二进制后出现的1的次数
func Test_CountOne(t *testing.T) {
	var countOne = func(i int) int {
		if i < 0 {
			i = -i
		}
		count := 0
		for i > 0 {
			if i&1 == 1 {
				count++
			}
			i = i >> 1
		}

		return count
	}

	fmt.Println("1:", countOne(1))
	fmt.Println("10:", countOne(10))
	fmt.Println("121:", countOne(121))
	fmt.Println("-1231:", countOne(-1231))
	fmt.Println("111:", countOne(111))
	fmt.Println("2:", countOne(2))
	fmt.Println("0:", countOne(0))
}

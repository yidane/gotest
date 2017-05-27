package testing

import "testing"
import "strings"

//测试单个文件 go test -v testing_test.go
//测试单个方法 go test -v testing_test.go -run Test_Index
func Test_Index(t *testing.T) {
	const s, sep, want = "chicken", "ken", 4
	got := strings.Index(s, sep)
	if got != want {
		t.Errorf("Index(%q,%q)=%v;want %v", s, sep, got, want)
	}
}

//表驱动测试
func Test_Index1(t *testing.T) {
	var tests = []struct {
		s   string
		sep string
		out int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"fo", "foo", -1},
		{"oofofoofooo", "f", 2},
	}
	for _, test := range tests {
		actual := strings.Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index (%q,%q) = %v;want %v", test.s, test.sep, actual, test.out)
		}
	}
}

/*
*testing.T参数用于错误报告：

t.Errorf("got bar = %v, want %v", got, want)
t.Fatalf("Frobnicate(%v) returned error: %v", arg, err)
t.Logf("iteration %v", i)

也可以用于enable并行测试(parallet test)：
t.Parallel()

控制一个测试是否运行：

if runtime.GOARCH == "arm" {
    t.Skip("this doesn't work on ARM")
}

我们用go test命令来运行特定包的测试。

默认执行当前路径下包的测试代码。

标准库的测试：
$ go test std
*/

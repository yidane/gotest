package benchmark

import "testing"

func Test_T(t *testing.T) {
	Test("http://localhost:8888/GET/", 3000)
}

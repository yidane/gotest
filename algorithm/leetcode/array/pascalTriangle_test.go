package array

import (
	"fmt"
	"testing"
)

func Test_pascalTriangle(t *testing.T) {
	result := pascalTriangle(10)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	t.Error(1)
}

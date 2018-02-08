/*
有个直角三角形三边为A,B,C三个整数。已知C为最长边。现告诉你A的长度，求一组B,C，使得B和C最接近

a^2+b^2=c^2

(2n+1)^2=(c+b)(c-b)    -- c-b=1
(2n+1)^2=2b+1
b=(a^2-1)/2
b=(a+1)(a-1)/2         -- 当a为奇数时,b为整数,c=b+1为整数

a^2=(c+b)(c-b)         -- c-b=2
a^2=(b+2+b)*2
a^2=4b+4
b=(a^2-4)/4
b=(a/2+1)(a/2-1)       -- 当a为偶数时，b为整数，c=b+2为整数

*/

package main

import (
	"fmt"
	"testing"
)

func Test_rightTriangle(t *testing.T) {
	for a := 1; a < 1000; a++ {
		b, c := rightTriangle(a)

		if b == 0 {
			fmt.Println(a)
		} else {
			fmt.Println(a, "^2 + ", b, "^2 = ", c, "^2", "	", c-b)
		}
	}

	t.Error(1)
}

package leetcode

import (
	"math"
	"strconv"
	"strings"
)

//*************************************************************************
/*
https://leetcode.com/problems/reverse-integer/#/description

Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

click to show spoilers.

Note:
The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.

Have you thought about this?
Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!
If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.
Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer, then the reverse of 1000000003 overflows. How should you handle such cases?
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
//*************************************************************************

func reverse(x int) int {
	out := ""
	f := false
	if x < 0 {
		f = true
		x = -x
	}
	for x > 0 {
		b := x % 10
		if b > 0 || (b == 0 && x > 0) {
			out += strconv.Itoa(x % 10)
		}

		x = x / 10
	}

	o, err := strconv.ParseInt(strings.Trim(out, "0"), 0, 0)
	if err != nil {
		return 0
	}

	if o > math.MaxInt32 {
		return 0
	}

	if o < math.MinInt32 {
		return 0
	}

	r := int(o)
	if f {
		r = -r
	}

	return r
}

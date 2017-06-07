package leetcode

import (
	"math"
	"strings"
)

//*************************************************************************
/*
https://leetcode.com/problems/string-to-integer-atoi/#/description

Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.
Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

Requirements for atoi:
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.
The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.
*/
//*************************************************************************

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}

	var result int64
	m := false
	f := false
	n := true
	for i := 0; i < len(str); i++ {
		var b int64
		switch str[i] {
		case '+':
			if m {
				return 0
			}
			m = true
			n = false
		case '-':
			if m {
				return 0
			}
			m = true
			f = true
			n = false
		case '1':
			b = 1
		case '2':
			b = 2
		case '3':
			b = 3
		case '4':
			b = 4
		case '5':
			b = 5
		case '6':
			b = 6
		case '7':
			b = 7
		case '8':
			b = 8
		case '9':
			b = 9
		case '0':
			b = 0
		default:
			i = len(str) + 1
			n = false
		}

		if n {
			result = result*10 + b
			if result > math.MaxInt32 {
				if f {
					return math.MinInt32
				}
				return math.MaxInt32
			}
		}
		n = true
	}

	if f {
		result = -result
	}

	return int(result)
}

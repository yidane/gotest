package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/powx-n/#/description

Implement pow(x, n).
*/
//*************************************************************************

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}

	var result float64 = 1
	for i := 0; i < n; i++ {
		result *= x
	}

	return result
}

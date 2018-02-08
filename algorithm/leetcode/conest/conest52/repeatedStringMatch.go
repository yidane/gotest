package conest52

import (
	"strings"
)

/*
686. Repeated String Match My SubmissionsBack to Contest
Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it. If no such solution, return -1.

For example, with A = "abcd" and B = "cdabcdab".

Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.
*/

func repeatedStringMatch(A string, B string) int {
	if len(A) == 0 {
		return -1
	}

	times := 1
	for len(A) < 10000 && strings.Index(A, B) == -1 {
		A += A
		times++
	}

	return times
}

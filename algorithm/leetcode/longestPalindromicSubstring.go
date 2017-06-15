package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/longest-palindromic-substring/#/description

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example:
Input: "cbbd"
Output: "bb"
*/
//*************************************************************************

type TwoPoint struct {
	Min int
	Max int
	Sub int
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	strArr := make(map[rune]*TwoPoint)
	for i, v := range s {
		if t, ok := strArr[v]; ok {
			if t.Min == -1 {
				t.Min = t.Max
			}
			t.Max = i
			t.Sub = t.Max - t.Min
		} else {
			strArr[v] = &TwoPoint{Max: i, Min: -1}
		}
	}

	two := TwoPoint{}
	for _, t := range strArr {
		if t.Sub > two.Sub {
			two.Max = t.Max
			two.Min = t.Min
			two.Sub = two.Max - two.Min
		}
	}

	if two.Sub > 0 {
		return s[two.Min : two.Max+1]
	}
	return ""
}

package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description

Given a string, find the length of the longest substring without repeating characters.

Examples:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
//*************************************************************************

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	bArr := make(map[rune]int)
	maxLength := 0
	begin := 0
	for i, v := range s {
		if j, ok := bArr[v]; ok {
			if i-begin > maxLength {
				maxLength = i - begin
			}
			if j+1 > begin {
				begin = j + 1
			}
		}
		bArr[v] = i
	}

	if len(s)-begin > maxLength {
		maxLength = len(s) - begin
	}

	return maxLength
}

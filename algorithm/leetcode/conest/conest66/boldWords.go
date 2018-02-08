package conest66

import (
	"sort"
)

/*
758. Bold Words in String My SubmissionsBack to Contest

Given a set of keywords words and a string S, make all appearances of all keywords in S bold. Any letters between <b> and </b> tags become bold.

The returned string should use the least number of tags possible, and of course the tags should form a valid combination.

For example, given that words = ["ab", "bc"] and S = "aabcd", we should return "a<b>abc</b>d". Note that returning "a<b>a<b>b</b>c</b>d" would use more tags, so it is incorrect.

Note:

words has length in range [0, 50].
words[i] has length in range [1, 10].
S has length in range [0, 500].
All characters in words[i] and S are lowercase letters.
*/

func boldWords(words []string, S string) string {
	if len(S) == 0 {
		return ``
	}
	boldIndexArr := []int{}
	Ss := []rune(S)
	for _, s := range words {
		if len(s) == 0 || len(s) > len(S) {
			continue
		}

	}

	sort.Ints(boldIndexArr)

	return ``
}

func findMaps(rs []rune, rs2 []rune, benginIndex int) []int {
	if len(rs2) == 0 {
		return []int{}
	}
	if len(rs)-benginIndex < len(rs2) {
		return []int{}
	}

	rtnIntArr := []int{}
	for i := benginIndex; i < len(rs); i++ {
		if len(rs)-benginIndex < len(rs2) {
			break
		}

		if rs[benginIndex] == rs2[0] {
			isMatch := true
			for j := 1; j < len(rs2); j++ {
				if rs[benginIndex+j] != rs2[j] {
					isMatch = false
				}
			}
			if isMatch {
				for i := 0; i < len(rs2); i++ {
					rtnIntArr = append(rtnIntArr, benginIndex+i)
				}
				benginIndex++
			} else {
				benginIndex += len(rs2)
			}
		} else {
			benginIndex += len(rs2)
		}
	}

	return rtnIntArr
}

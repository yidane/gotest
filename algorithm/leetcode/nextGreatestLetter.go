package leetcode

/*
Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.

Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.

Examples:
Input:
letters = ["c", "f", "j"]
target = "a"
Output: "c"

Input:
letters = ["c", "f", "j"]
target = "c"
Output: "f"

Input:
letters = ["c", "f", "j"]
target = "d"
Output: "f"

Input:
letters = ["c", "f", "j"]
target = "g"
Output: "j"

Input:
letters = ["c", "f", "j"]
target = "j"
Output: "c"

Input:
letters = ["c", "f", "j"]
target = "k"
Output: "c"
Note:
1.letters has a length in range [2, 10000].
2.letters consists of lowercase letters, and contains at least 2 unique letters.
3.target is a lowercase letter.
*/

func nextGreatestLetter(letters []byte, target byte) byte {
	//直接遍历处理，算法复杂度O(n)
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}

func nextGreatestLetter2(letters []byte, target byte) byte {
	//采用2分法处理，算法复杂度O(log(n))
	l := len(letters)
	i, j := 0, l
	for l > i {
		j = i + (l-i)/2
		switch {
		case letters[j] > target:
			l = j
		case letters[j] <= target:
			i = j + 1
		}
	}
	return letters[j%l]
}

// func nextGreatestLetter2(letters []byte, target byte) byte {
// 	n := len(letters)

// 	//hi starts at 'n' rather than the usual 'n - 1'.
// 	//It is because the terminal condition is 'lo < hi' and if hi starts from 'n - 1',
// 	//we can never consider value at index 'n - 1'
// 	lo, hi := 0, n

// 	//Terminal condition is 'lo < hi', to avoid infinite loop when target is smaller than the first element
// 	for lo < hi {
// 		mid := lo + (hi-lo)/2
// 		if letters[mid] > target {
// 			hi = mid
// 		} else {
// 			lo = mid + 1
// 		} //a[mid] <= x
// 	}

// 	//Because lo can end up pointing to index 'n', in which case we return the first element
// 	return letters[lo%n]
// }

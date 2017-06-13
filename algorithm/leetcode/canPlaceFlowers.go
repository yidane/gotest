package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/can-place-flowers/#/description

Suppose you have a long flowerbed in which some of the plots are planted and some are not.
However, flowers cannot be planted in adjacent plots - they would compete for water and both would die.
Given a flowerbed (represented as an array containing 0 and 1, where 0 means empty and 1 means not empty), and a number n,
return if n new flowers can be planted in it without violating the no-adjacent-flowers rule.

Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: True
Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: False
Note:
The input array won't violate no-adjacent-flowers rule.
The input array size is in the range of [1, 20000].
n is a non-negative integer which won't exceed the input array size.
*/
//*************************************************************************
func canPlaceFlowers0(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	l := len(flowerbed)
	if l < n || l == 0 {
		return false
	}
	if l == 1 {
		if flowerbed[0] == 1 {
			return false
		}
		return n <= 1
	}
	if l == 2 {
		if flowerbed[0] == 1 || flowerbed[1] == 1 {
			return false
		}
		return n <= 1
	}

	blankArr := []int{}
	blank := 0
	if flowerbed[0] == 0 {
		blank = 2
	}

	for i := 1; i < l-1; i++ {
		if flowerbed[i] == 0 {
			blank++
		} else {
			if blank >= 3 {
				blankArr = append(blankArr, blank)
			}
			blank = 0
		}
	}

	if flowerbed[l-1] == 0 && blank >= 1 {
		blank += 2
	}
	if blank >= 3 {
		blankArr = append(blankArr, blank)
	}

	total := 0
	for _, v := range blankArr {
		total += (v - 1) / 2
	}

	return total >= n
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	for i := 0; i < l && n > 0; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}
		if i < l-1 && flowerbed[i+1] == 1 {
			continue
		}
		n--
		flowerbed[i] = 1
	}

	return n == 0
}

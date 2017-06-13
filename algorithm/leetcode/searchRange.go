package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/search-for-a-range/#/description

Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

For example,
Given [5, 7, 7, 8, 8, 10] and target value 8,
return [3, 4].
*/
//*************************************************************************

func searchRange(nums []int, target int) []int {
	begin := -1
	end := -1
	for i := 0; i < len(nums); i++ {
		r := len(nums) - i - 1
		if r >= i {
			if nums[i] > target {
				break
			}
			if nums[r] < target {
				break
			}
			if nums[i] == target {
				if begin == -1 || begin > i {
					begin = i
				}
				if end == -1 || end < i {
					end = i
				}
			}
			if nums[r] == target {
				if end == -1 || end < r {
					end = r
				}
				if begin == -1 || begin > r {
					begin = r
				}
			}
		} else {
			break
		}
	}

	return []int{begin, end}
}

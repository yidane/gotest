package leetcode

import (
	"fmt"
)

//*************************************************************************
/*
https://leetcode.com/problems/two-sum/#/description

Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
//*************************************************************************

//N!
func twoSum0(nums []int, target int) []int {
	if nums == nil || len(nums) <= 1 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println(nums[i], nums[j])
				return []int{i, j}
			}
		}
	}

	return nil
}

//2n
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) <= 1 {
		return nil
	}

	arr := make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if ii, ok := arr[nums[i]]; ok {
			ii = append(ii, i)
		} else {
			arr[nums[i]] = []int{i}
		}
	}

	for v, ii := range arr {
		nv := target - v
		if v == nv {
			l := len(ii)
			if l >= 2 {
				return []int{ii[0], ii[1]}
			}
			continue
		}
		if jj, ok := arr[nv]; ok {
			return []int{ii[0], jj[0]}
		}
	}

	return nil
}

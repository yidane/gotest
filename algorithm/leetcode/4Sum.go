package leetcode

import (
	"fmt"
	"sort"
)

/*
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note: The solution set must not contain duplicate quadruplets.

For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	resultMap := map[string][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > 0; j-- {
			if i < j {
				k := i + 1
				m := j - 1
				total := 0
				for k < m {
					total = nums[i] + nums[k] + nums[m] + nums[j]
					switch {
					case total < target:
						k++
					case total == target:
						key := fmt.Sprintf("%d_%d_%d_%d", nums[i], nums[k], nums[m], nums[j])
						if _, ok := resultMap[key]; !ok {
							resultMap[key] = []int{nums[i], nums[k], nums[m], nums[j]}
						}
						k++
					case total > target:
						m--
					}
				}
			}
		}
	}

	var result [][]int

	for _, v := range resultMap {
		result = append(result, v)
	}

	return result
}

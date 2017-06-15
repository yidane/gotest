package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/3sum/#/description

Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
//*************************************************************************

func threeSum0(nums []int) [][]int {
	result := [][]int{}
	mapKey := make(map[int]bool)
	if len(nums) > 2 {
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				for k := j + 1; k < len(nums); k++ {
					if nums[i]+nums[j]+nums[k] == 0 {
						ii := nums[i]
						jj := nums[j]
						kk := nums[k]
						if ii < jj {
							ii, jj = jj, ii
						}
						if ii < kk {
							ii, kk = kk, ii
						}
						if jj < kk {
							jj, kk = kk, jj
						}
						if ii == jj || jj == kk || ii == kk {
							if _, ok := mapKey[jj]; !ok {
								mapKey[jj] = true
								result = append(result, []int{nums[i], nums[j], nums[k]})
							}
						}
					}
				}
			}
		}
	}

	return result
}

//复杂度:O(n^2)，采用双向循环方式，妙！
func threeSum(nums []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			l := i + 1
			r := len(nums) - 1
			sum := -nums[i]
			for l < r {
				if nums[l]+nums[r] == sum {
					result = append(result, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[l]+nums[r] > sum {
					r--
				} else {
					l++
				}
			}
		}
	}

	return result
}

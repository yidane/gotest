package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/median-of-two-sorted-arrays/#/description

There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
//*************************************************************************

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	arr := make(map[int]int, l)

	i, j := 0, 0
	for i+j < l {
		if nums1[i] < nums2[j] {
			arr[i+j] = nums1[i]
			if i < l1 {
				i++
			} else {
				if j < l2 {
					j++
				}
			}
		} else {
			arr[i+j] = nums1[j]
			if j < l2 {
				j++
			} else {
				if i < l1 {
					i++
				}
			}
		}
	}

	if l%2 == 0 {
		return (float64)(arr[l/2]+arr[l/2+1]) / 2
	}
	return (float64)(arr[l/2+1])
}

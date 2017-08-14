package mock

func findKthLargest(nums []int, k int) int {
	if len(nums) < k || k < 0 {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j >= i; j-- {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums[len(nums)-k]
}

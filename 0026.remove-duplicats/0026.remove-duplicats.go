package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nextIndex := 1 // first item never changes, also this means length
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nextIndex] = nums[i]
			nextIndex++
		}
	}
	return nextIndex
}

// RemoveDuplicates export for test
func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}

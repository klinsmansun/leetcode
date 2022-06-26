package medianoftwosortedarrays

func getNextNumberFunc(nums1 []int, nums2 []int) func() (result int) {
	index1, index2 := 0, 0

	return func() (result int) {
		// return num2[index] in one of the 2 conditions
		// 1. num1 has no more items
		// 2. (when 1 failed) num2 still has item and nums[index2] < nums1[index1]
		if index1 >= len(nums1) || (index2 < len(nums2) && nums2[index2] < nums1[index1]) {
			result = nums2[index2]
			index2++
		} else {
			result = nums1[index1]
			index1++
		}
		return
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// getNextNumberFunc of nums1 and nums2
	gn := getNextNumberFunc(nums1, nums2)
	if (len(nums1)+len(nums2))%2 == 1 {
		// we want the index n/2+1
		for i := 0; i < (len(nums1)+len(nums2))/2; i++ {
			gn() // don't need it, skip
		}
		return float64(gn()) // return target
	}

	// we want the average of n/2-1 and n/2+1
	for i := 0; i < (len(nums1)+len(nums2))/2-1; i++ {
		gn()
	}
	return (float64(gn()) + float64(gn())) / 2
}

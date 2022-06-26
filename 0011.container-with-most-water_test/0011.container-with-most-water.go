package leetcode

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// length of input is at least 2
// ! key idea is first and last have biggest x
// ! if x decrease and area increase, means a higher y is necessasy
func maxArea(height []int) int {
	left, right, result := 0, len(height)-1, 0
	//baseHeight := 0
	for left < right {
		//baseHeight = min(height[left], height[right])
		result = max(result, (right-left)*min(height[left], height[right]))
		// if we want to increase left or right multiple times, we need one more variable (to keep current lower height)
		// it will slower (since one additional memory allocation)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	//left = baseHeight
	return result
}

// ContainerWithMostWater is used to expose maxArea to test
func ContainerWithMostWater(height []int) int {
	return maxArea(height)
}

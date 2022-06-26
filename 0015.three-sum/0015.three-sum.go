package leetcode

import (
	"sort"
)

// we eliminate duplicated by making return values are ordered by array index
func twoSumAdv(inputMap map[int]int, targetValue int) [][]int {
	result := make([][]int, 0)
	// find combination for each key
	for key, totalCount := range inputMap {
		targetKey := targetValue - key
		// key must be bigger than first digit (0-targetValue)
		if 0-targetValue > key || key > targetKey { // only use small key to find big key, so skip this situation
			continue
		} else if key == targetKey {
			if totalCount >= 2 {
				// 2 digits are the same, check digit appearence
				result = append(result, []int{key, key})
			}
			continue
		}
		if _, ok := inputMap[targetKey]; ok {
			result = append(result, []int{key, targetKey})
		}
	}
	return result
}

func threeSumNotGood(nums []int) [][]int {
	result := make([][]int, 0)
	inputMap := make(map[int]int) // { arrayValue: numbers-found }
	sort.Ints(nums)
	currentStart := 1 // never happens, so used as initial value

	// handle [0, 0, 0] at last, since
	// 1. it's the only case that first and third digit is the same (3 the same digits)
	// 2. it's the only case starting with 0

	for _, value := range nums {
		inputMap[value]++
	}

	for _, value := range nums {
		if value >= 0 { // never starting with positive numbers and handles zero at last
			break
		}
		if value == currentStart { // already processed
			continue
		}
		currentStart = value
		twoResult := twoSumAdv(inputMap, 0-value)

		for _, retValue := range twoResult {
			// the first 2 digit is the same, so we need to check if it appears wtice
			if value == retValue[0] && inputMap[value] < 2 {
				continue
			}
			result = append(result, []int{value, retValue[0], retValue[1]})
		}
	}

	if inputMap[0] >= 3 {
		result = append(result, []int{0, 0, 0})
	}
	// currentStart := 1 // never happens, so used as initial value
	return result
}

// revised
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	inputLength := len(nums)
	result := make([][]int, 0)

	for i := 0; i < inputLength-2; i++ {
		startIdx, secondIdx, thirdIdx := i, i+1, inputLength-1
		// it's impossible to start with number > 0
		if nums[startIdx] > 0 {
			return result
		}
		// the same starting number checked
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		targetSum := 0 - nums[startIdx] // target sum for second and third
		// find all result for starting digit nums[startIdx]
		for secondIdx < thirdIdx {
			if nums[secondIdx]+nums[thirdIdx] == targetSum {
				// found a match result
				result = append(result, []int{nums[startIdx], nums[secondIdx], nums[thirdIdx]})
				// increase second and decrease third to find next possible match
				for secondIdx < thirdIdx {
					secondIdx++
					if nums[secondIdx] != nums[secondIdx-1] {
						break
					}
				}
				for secondIdx < thirdIdx {
					thirdIdx--
					if nums[thirdIdx] != nums[thirdIdx+1] {
						break
					}
				}
				// here is the trick, sum too big, decrease third
			} else if nums[secondIdx]+nums[thirdIdx] > targetSum {
				for secondIdx < thirdIdx {
					thirdIdx--
					if nums[thirdIdx] != nums[thirdIdx+1] {
						break
					}
				}
				// here is the trick, sum too small increase second
			} else { // nums[secondIdx]+nums[thirdIdx] < targetSum
				for secondIdx < thirdIdx {
					secondIdx++
					if nums[secondIdx] != nums[secondIdx-1] {
						break
					}
				}
			}
		}
	}

	return result
}

// ThreeSum export for test
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

// ThreeSumNotGood export for test
func ThreeSumNotGood(nums []int) [][]int {
	return threeSumNotGood(nums)
}

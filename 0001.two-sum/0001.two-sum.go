package twosum

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	pool := make(map[int]int) // { candidateValue: candidateIndex }
	for i := 0; i < len(nums); i++ {
		if candidateIndex, found := pool[target-nums[i]]; found { // candidateValue := target-nums[i]
			ret[0] = candidateIndex
			ret[1] = i
			return ret
		}
		pool[nums[i]] = i // not found, add current value and index to pool
	}
	return ret
}

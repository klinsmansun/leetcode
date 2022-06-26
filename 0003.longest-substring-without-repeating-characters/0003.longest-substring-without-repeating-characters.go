package leetcode

func lengthOfLongestSubstring(s string) int {
	uArray := []uint8(s)
	length, left, right, maxLength := len(uArray), 0, 1, 1

	if length < 2 {
		return length
	}

	candidateSlice := uArray[:1]

	for right < len(uArray) {
		found := -1
		for index, value := range candidateSlice {
			if value == uArray[right] { // found repeating
				found = index
				break
			}
		}
		if found > -1 { // found repeating, check current candidate
			if len(candidateSlice) > maxLength {
				maxLength = len(candidateSlice)
			}
			// start from next letter after found one
			// ! starting with letter between left and found can be skipped (must be shorter than starting with current left)
			left = left + found + 1
		}
		right++                             // next character to check
		candidateSlice = uArray[left:right] // update candidate index
	}

	// string ends, checks if final result longer than maxLength
	if len(candidateSlice) > maxLength {
		maxLength = len(candidateSlice)
	}

	return maxLength
}

// LengthOfLongestSubstring exports for test
func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

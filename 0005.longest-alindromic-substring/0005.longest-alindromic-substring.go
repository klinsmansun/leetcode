package leetcode

func longestPalindrome(s string) string {
	charArray := []uint8(s)
	start := 0 // start index
	arrayLength := len(charArray)
	maxLength := 1 // current max found
	// empty string
	if arrayLength == 0 {
		return s
	}

	// reverse compare for each character
	for i := 0; i < arrayLength; i++ {
		posMatch := i
		revMatch := i - 1
		leng := 0
		tempStart := 0

		// first character
		if revMatch < 0 {
			continue // continue to next character
		}

		// situation 1: start comparing with repeated character
		for {
			// index out of range
			if revMatch < 0 || posMatch == arrayLength {
				break
			}
			if charArray[revMatch] == charArray[posMatch] {
				leng += 2
				tempStart = revMatch
				posMatch++
				revMatch--
				continue
			}
			break // not match, ends here
		}
		if leng > maxLength {
			start = tempStart
			maxLength = leng
		}

		posMatch = i
		revMatch = i - 2 // start with the left x 2 character
		leng = 1         // set to 1, if match, we can +2
		tempStart = 0

		// first 2 characters
		if revMatch < 0 {
			continue // continue to next character
		}

		// situation 2: start comparing with left x 2 character
		for {
			// index out of range
			if revMatch < 0 || posMatch == arrayLength {
				break
			}
			if charArray[revMatch] == charArray[posMatch] {
				leng += 2
				tempStart = revMatch
				posMatch++
				revMatch--
				continue
			}
			break // not match, ends here
		}
		if leng > maxLength {
			start = tempStart
			maxLength = leng
		}
	}
	return string(charArray[start : start+maxLength])
}

// LongestPalindrome exports for test
func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

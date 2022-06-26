package integertoroman

// I  1
// V  5
// X  10
// L  50
// C  100
// D  500
// M  1000

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

func handleFour(upperBound, lowerBound, remain int, candidate string) (left int, result string) {
	if remain >= lowerBound && remain < upperBound {
		left = remain - lowerBound
		result = candidate
	} else {
		left = remain
		result = ""
	}

	return
}

func handleNine(boundary, remain int, candidate string) (left int, result string) {
	if remain >= boundary {
		left = remain - boundary
		result = candidate
	} else {
		left = remain
		result = ""
	}

	return
}

func handleDec(boundary, remain int, candidate string) (left int, result string) {
	for i := 0; i < remain/boundary; i++ {
		result += candidate
	}
	left = remain % boundary

	return
}

func intToRoman(num int) string {
	append := ""

	// 1000
	remain, result := handleDec(1000, num, "M")

	// 900
	remain, append = handleNine(900, remain, "CM")
	result = result + append

	// 500
	if remain >= 500 {
		result += "D"
		remain -= 500
	}

	// 400
	remain, append = handleFour(500, 400, remain, "CD")
	result = result + append

	// 100
	remain, append = handleDec(100, remain, "C")
	result += append

	// 90
	remain, append = handleNine(90, remain, "XC")
	result = result + append

	// 50
	if remain >= 50 {
		result += "L"
		remain -= 50
	}

	// 40
	remain, append = handleFour(50, 40, remain, "XL")
	result = result + append

	// 10
	remain, append = handleDec(10, remain, "X")
	result += append

	// 9
	remain, append = handleNine(9, remain, "IX")
	result = result + append

	// 5
	if remain >= 5 {
		result += "V"
		remain -= 5
	}

	// 4
	remain, append = handleFour(5, 4, remain, "IV")
	result = result + append

	// 1
	remain, append = handleDec(1, remain, "I")
	result += append

	return result
}

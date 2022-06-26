package leetcode

func reverse(x int) int {
	// note that range of 32bit int is -2^31 ~ 2^31-1, -2147483648 ~ 2147483647
	opCode := 1 // handle positive/negative
	opNumber := x
	newNumber := 0
	if opNumber < 0 {
		opCode = -1
		opNumber = x * (-1)
	}
	// first 9 times never overflows
	for i := 0; i < 9; i++ {
		lastDigit, left := opNumber%10, opNumber/10
		opNumber = left
		newNumber = newNumber*10 + lastDigit
		if opNumber == 0 { // no more digits
			return newNumber * opCode
		}
	}

	// if we reach here, opNumber must be 1 or 2(leading digit of input)
	// last digit, check if we will overflow
	// if numNumber == 214748364, it never overflows with opNumber = 1 or 2, no need to check
	// just check newNumber > 214748364
	if newNumber > 214748364 {
		return 0
	}
	return (newNumber*10 + opNumber) * opCode
}

// ReverseInteger proxies reverse() to test
func ReverseInteger(x int) int {
	return reverse(x)
}

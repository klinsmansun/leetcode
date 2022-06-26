package romantointeger

var romanMapping map[byte]int = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

// reference 0012 for roman definition
func romanToInt(s string) int {
	ca := []byte(s)
	result := 0

	for i := 0; i < len(ca); i++ {
		if i+1 < len(ca) && romanMapping[ca[i]] < romanMapping[ca[i+1]] {
			result -= romanMapping[ca[i]]
		} else {
			result += romanMapping[ca[i]]
		}
	}

	return result
}

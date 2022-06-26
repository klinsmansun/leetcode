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

type checkPair struct {
	roman    string
	boundary int
}

var checkList []checkPair = []checkPair{
	{roman: "M", boundary: 1000},
	{roman: "CM", boundary: 900},
	{roman: "D", boundary: 500},
	{roman: "CD", boundary: 400},
	{roman: "C", boundary: 100},
	{roman: "XC", boundary: 90},
	{roman: "L", boundary: 50},
	{roman: "XL", boundary: 40},
	{roman: "X", boundary: 10},
	{roman: "IX", boundary: 9},
	{roman: "V", boundary: 5},
	{roman: "IV", boundary: 4},
	{roman: "I", boundary: 1},
}

func intToRoman2(num int) string {
	result := ""
	remain := num
	for _, check := range checkList {
		for remain >= check.boundary {
			result += check.roman
			remain -= check.boundary
		}
	}

	return result
}

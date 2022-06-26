package leetcode

import (
	"fmt"
	"testing"

	"klinsman.csie.org/klinsman/lib"
)

var reverseIntegerTestData = map[int]int{
	123:         321,
	-123:        -321,
	120:         21,
	1463847412:  2147483641,
	-1463847412: -2147483641,
	1463847413:  0,
	1563847412:  0,
	1234567809:  0,
}

func TestReverseInterger(t *testing.T) {
	pass, fail, skip, all := 0, 0, 0, 0
	for key, value := range reverseIntegerTestData {
		all++
		if result := ReverseInteger(key); result != value {
			fail++
			t.Error(lib.PaintString("red", fmt.Sprintln("Test Failed, expect:", value, "got:", result)))
			continue
		}
		pass++
	}

	t.Log(lib.FormatTestStatistic(all, pass, fail, skip))
}

func BenchmarkReverseInterger(b *testing.B) {
	for key, value := range reverseIntegerTestData {
		if result := ReverseInteger(key); result != value {
			ReverseInteger(key)
		}
	}
}

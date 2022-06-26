package leetcode

import (
	"testing"
)

func Test3sum(t *testing.T) {
	result := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	result2 := ThreeSum([]int{3, 0, -2, -1, 1, 2})
	if len(result) != 2 {
		t.Fail()
	}
	if len(result2) != 3 {
		t.Fail()
	}
}

func Test3sumNotGood(t *testing.T) {
	result := ThreeSumNotGood([]int{-1, 0, 1, 2, -1, -4})
	result2 := ThreeSumNotGood([]int{3, 0, -2, -1, 1, 2})
	if len(result) != 2 {
		t.Fail()
	}
	if len(result2) != 3 {
		t.Fail()
	}
}

var threesumTestData []int

func init() {
	threesumTestData = make([]int, 10000)
}

func Benchmark3sum(b *testing.B) {
	for i := 0; i < 1000; i++ {
		ThreeSum(threesumTestData)
	}
}

func Benchmark3sumNotGood(b *testing.B) {
	for i := 0; i < 1000; i++ {
		ThreeSumNotGood(threesumTestData)
	}
}

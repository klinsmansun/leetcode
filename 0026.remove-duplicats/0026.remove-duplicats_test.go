package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// assume arrays are sorted
	test1 := []int{1, 1, 2}
	test2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	test3 := make([]int, 0)
	count1 := RemoveDuplicates(test1)
	count2 := RemoveDuplicates(test2)
	count3 := RemoveDuplicates(test3)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	if count1 != 2 {
		t.Fail()
	}
	if count2 != 5 {
		t.Fail()
	}
	if count3 != 0 {
		t.Fail()
	}
}

package leetcode

import (
	"testing"
)

var testData []*ListNode
var testDataEmpty []*ListNode
var expect []int

func TestMergeKList(t *testing.T) {
	result := MergeKLists(testData)
	current := result
	for i := 0; i < len(expect); i++ {
		if expect[i] != current.Val {
			t.Fail()
		}
		current = current.Next
	}

	result = MergeKLists(testDataEmpty)
	if result != nil {
		t.Fail()
	}
}

func insertTestData(index int, data []int) {
	var current *ListNode
	for key, value := range data {
		if key == 0 {
			testData[index] = &ListNode{Val: value, Next: nil}
			current = testData[index]
		} else {
			current.Next = &ListNode{Val: value, Next: nil}
			current = current.Next
		}
	}
}

func init() {
	// leetcode test case 1
	test1 := []int{1, 4, 5}
	test2 := []int{1, 3, 4}
	test3 := []int{2, 6}
	expect = []int{1, 1, 2, 3, 4, 4, 5, 6}
	testData = make([]*ListNode, 3)
	insertTestData(0, test1)
	insertTestData(1, test2)
	insertTestData(2, test3)

	// leetcode test case 2, 2 empty lists
	testDataEmpty = make([]*ListNode, 2)
}

package addtwonumbers

import (
	"fmt"
	"reflect"
	"testing"
)

func createNode(n int) (left int, pNewNode *ListNode) {
	pNewNode = &ListNode{Val: n % 10}
	return n / 10, pNewNode
}

// n is expected to be a positive number
func createTestNodeList(n int) *ListNode {
	// create first node
	left, result := createNode(n)
	pHead := result

	// create other nodes
	for left > 0 {
		left, pHead.Next = createNode(left)
		pHead = pHead.Next
	}

	return result
}

func printNodeList(input *ListNode) {
	head := input
	for {
		fmt.Printf("%v ", head.Val)
		if head.Next != nil {
			fmt.Printf("-> ")
			head = head.Next
		} else {
			fmt.Printf("\n")
			break
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "default case",
			args: args{
				l1: createTestNodeList(342),
				l2: createTestNodeList(465),
			},
			want: createTestNodeList(807),
		},
		{
			name: "failed case 1",
			args: args{
				l1: createTestNodeList(0),
				l2: createTestNodeList(0),
			},
			want: createTestNodeList(0),
		},
		{
			name: "failed case 2, contains 0 in the middle",
			args: args{
				l1: createTestNodeList(1207),
				l2: createTestNodeList(1201),
			},
			want: createTestNodeList(2408),
		},
		{
			name: "other case 1, different length with 9",
			args: args{
				l1: createTestNodeList(9999999),
				l2: createTestNodeList(9),
			},
			want: createTestNodeList(10000008),
		},
		{
			name: "other case 2, different length with 0",
			args: args{
				l1: createTestNodeList(9999999),
				l2: createTestNodeList(0),
			},
			want: createTestNodeList(9999999),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printNodeList(tt.want)
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

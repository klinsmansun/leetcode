package leetcode

// ListNode from leetcode
/** Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// brute force, bad performance
// this function finds the smallest item through all lists
// so we need call this multiple times until lists all empty
func FindSmallest(lists []*ListNode) *ListNode {
	var result *ListNode = nil
	var found int
	for index, pListHead := range lists {
		switch {
		case pListHead == nil:
			continue
		case result == nil:
			result = pListHead
			found = index
		case pListHead.Val < result.Val:
			result = pListHead
			found = index
		}
	}
	if result != nil {
		lists[found] = lists[found].Next
	}

	return result
}

func merge2Lists(first *ListNode, second *ListNode) *ListNode {
	// input may be empty
	// if one is null, return another (no matter if another one is nil or not)
	if first == nil {
		return second
	} else if second == nil {
		return first
	}

	// if we reach here, both ListNode are not nil
	// handle first one item (result is nil now)
	var result, current *ListNode
	if first.Val < second.Val {
		result = first
		current = first
		first = first.Next
	} else {
		result = second
		current = second
		second = second.Next
	}

	// do this until one of first/second becomes nil
	for first != nil && second != nil {
		if first.Val < second.Val {
			current.Next = first // use current to append list
			current = first
			first = first.Next
		} else {
			current.Next = second
			current = second
			second = second.Next
		}
	}

	// one is nil (or both), append rest items of the other list(nil is also ok)
	if first != nil {
		current.Next = first
	} else {
		current.Next = second
	}

	return result
}

// divide and conquer
func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return merge2Lists(lists[0], lists[1])
	default:
		return merge2Lists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
	}
}

// MergeKLists exports mergeKLists for tests
func MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

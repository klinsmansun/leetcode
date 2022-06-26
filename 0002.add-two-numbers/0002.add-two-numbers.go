package addtwonumbers

// ListNode defines node for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// no need to check exit condition, we check that in addTwoNumbers()
// c language style
func createResultNode(ppFirst **ListNode, ppSecond **ListNode, exceed *int, ppNext **ListNode) **ListNode {
	// calculate this digit
	if *ppFirst != nil {
		*exceed += (*ppFirst).Val
		*ppFirst = (*ppFirst).Next
	}
	if *ppSecond != nil {
		*exceed += (*ppSecond).Val
		*ppSecond = (*ppSecond).Next
	}

	// store digit and exceed
	*ppNext = &ListNode{Val: *exceed % 10}
	*exceed = *exceed / 10
	return &(*ppNext).Next // pp to next candidate
}

// use multiple return values instead of point-to-pointer
func createResultNodeLessPointer(pFirst, pSecond *ListNode, exceed int) (pFirstR, pSecondR, pNewNode *ListNode, exceedR int) {
	if pFirst != nil {
		exceed += pFirst.Val
		pFirstR = pFirst.Next
	}
	if pSecond != nil {
		exceed += pSecond.Val
		pSecondR = pSecond.Next
	}

	pNewNode = &ListNode{Val: exceed % 10}
	exceedR = exceed / 10

	return
}

// l1 and l2 are 2 linked list, each represents a integer
// each node is a 0-9 digit, stored into list in reverse order
// ex: 123 becomes 3 -> 2 -> 1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode

	// create first node
	first, second, result, exceed := createResultNodeLessPointer(l1, l2, 0)
	pHead := result

	// check if we still need to increase list
	for exceed > 0 || first != nil || second != nil {
		first, second, pHead.Next, exceed = createResultNodeLessPointer(first, second, exceed)
		pHead = pHead.Next
	}

	return result
}

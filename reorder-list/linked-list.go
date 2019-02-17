/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head, appendage *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	doneHead := appendage
	for todoHead := head; nil != todoHead; {
		todoHead.Next, doneHead, todoHead = doneHead, todoHead, todoHead.Next
	}

	return doneHead
}

func getMiddle(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	oneStep, twoStep := head, head.Next
	for ; nil != twoStep && nil != twoStep.Next; oneStep, twoStep = oneStep.Next, twoStep.Next.Next {}

	return oneStep
}

func reorderList(head *ListNode)  {
	if nil == head || nil == head.Next {
		return
	}

	middle := getMiddle(head)
	l1, l2 := head, reverseList(middle.Next, nil)
	middle.Next = nil

	dummyNode := ListNode{0, nil}
	tail := &dummyNode

	for ; nil != l2; {
		tail.Next, l1, tail = l1, l1.Next, l1
		tail.Next, l2, tail = l2, l2.Next, l2
	}
	if nil != l1 {
		tail.Next = l1
	}

	return
}
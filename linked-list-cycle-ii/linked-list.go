/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return nil
	}
	if head == head.Next {
		return head
	}

	dummyNode := ListNode{0, head}
	oneStep, twoStep := head, head.Next
	for ; oneStep != twoStep; {
		if nil == twoStep || nil == twoStep.Next {
			return nil
		}
		if head == twoStep || head == twoStep.Next {
			return head
		}

		oneStep, twoStep = oneStep.Next, twoStep.Next.Next
    }

	for tmp := &dummyNode; tmp != oneStep; oneStep, tmp = oneStep.Next, tmp.Next {}

	return oneStep
}
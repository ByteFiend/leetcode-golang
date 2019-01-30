/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if nil == head || n <= 0 {
		return head
	}

	tail := head
	for i := 0; i < n; i++ {
		if nil == tail {
			return head
		}
		tail = tail.Next
	}

	dummyNode := ListNode{0, head}
	prev := &dummyNode
	for ; nil != tail; tail = tail.Next {
		prev = prev.Next
	}

	prev.Next = prev.Next.Next
	return dummyNode.Next
}
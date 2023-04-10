package code

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var a, b = headA, headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

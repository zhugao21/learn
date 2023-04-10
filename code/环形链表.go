package code

func detectCycle(head *ListNode) *ListNode {
	var fast = head
	var slow = head

	var commonFirst *ListNode
	for fast != nil && slow != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		if fast == slow {
			commonFirst = fast
			break
		}
	}
	if commonFirst == nil {
		return nil
	}

	var third = head
	for third != commonFirst {
		third = third.Next
		commonFirst = commonFirst.Next
	}
	return third
}

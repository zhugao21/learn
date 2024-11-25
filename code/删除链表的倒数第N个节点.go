package code

// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var cur = head
	for cur != nil && n > 0 {
		cur = cur.Next
		n--
	}
	if n > 0 {
		return head
	}
	if cur == nil {
		return head.Next
	}

	var tmp = head
	var pre *ListNode
	for cur != nil {
		cur = cur.Next
		pre = tmp
		tmp = tmp.Next
	}
	pre.Next = tmp.Next
	return head
}

func removeNthFromEnd01(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
func removeNthFromEnd02(head *ListNode, n int) *ListNode {
	var cur = head
	for n > 1 {
		cur = cur.Next
		n--
	}

	var newHead = &ListNode{}
	newHead.Next = head
	var newCur = newHead

	for cur.Next != nil {
		cur = cur.Next
		newCur = newCur.Next
	}

	newCur.Next = newCur.Next.Next
	return newHead.Next
}

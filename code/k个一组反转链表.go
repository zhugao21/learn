package code

// 1-2-3-4-5 k=2
// 2-1 3-4-5
func reverseKGroup(head *ListNode, k int) *ListNode {
	var cur = head
	var newHead, tail *ListNode

	for cur != nil {
		tmp, tmpCur := reverseKList(cur, k)
		if newHead == nil {
			newHead = tmp
		}
		if tail == nil {
			tail = cur
		} else {
			tail.Next = tmp
			tail = cur
		}
		cur = tmpCur
	}

	return newHead
}

func reverseKList(head *ListNode, k int) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head

	var cp = head
	var count int
	for cp != nil && count < k {
		count++
		cp = cp.Next
	}
	if count < k {
		return head, nil
	}

	for cur != nil && k > 0 {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		k--
	}

	return pre, cur
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var cur = head
	var tail = head
	k := right - left + 1

	for cur != nil && left > 1 {
		tail = cur
		cur = cur.Next
		left--
	}
	if tail != nil {
		tmp, tmpCur := reverseKList(cur, k)
		tail.Next = tmp
		cur.Next = tmpCur
	}
	return head
}

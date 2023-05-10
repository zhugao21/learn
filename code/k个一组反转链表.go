package code

// 1-2-3-4-5 k=2
// 2-1 3-4-5
func reverseKGroup(head *ListNode, k int) *ListNode {
	var length int
	cur := head
	for cur != nil {
		cur = cur.Next
		length++
	}

	var newHead, tail *ListNode
	cur = head
	for i := 0; i < length/k; i++ {
		tmpHead, nextHead := reverseKList(cur, k)
		if newHead == nil {
			newHead = tmpHead
		}
		if tail != nil {
			tail.Next = tmpHead
		}
		tail = cur
		cur = nextHead
	}
	if cur != nil {
		tail.Next = cur
	}
	return newHead
}

func reverseKList(head *ListNode, k int) (*ListNode, *ListNode) {
	var pre *ListNode
	var cur = head
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

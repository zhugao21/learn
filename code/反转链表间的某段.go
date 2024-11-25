package code

// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
func reverseBetween01(head *ListNode, left int, right int) *ListNode {
	var pre *ListNode
	var cur = head
	var pos = 1

	var newHead = head

	for pos < left {
		pre = cur
		cur = cur.Next
		pos++
		right--
	}
	res, tail := reverseK(cur, right)
	if pre != nil {
		pre.Next = res
	} else {
		newHead = res
	}
	cur.Next = tail
	return newHead
}

func reverseK(head *ListNode, k int) (*ListNode, *ListNode) {
	var pre *ListNode
	var cur = head
	for k > 0 {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
		k--
	}
	return pre, cur
}

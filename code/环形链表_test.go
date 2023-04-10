package code

import (
	"testing"
)

func Test_detectCycle(t *testing.T) {
	a := &ListNode{Val: 3}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 0}
	d := &ListNode{Val: -4}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b
	t.Log(detectCycle(a))
}

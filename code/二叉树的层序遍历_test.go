package code

import (
	"testing"
)

func Test_levelOrder(t *testing.T) {
	a := &TreeNode{Val: 3}
	b := &TreeNode{Val: 9}
	c := &TreeNode{Val: 20}
	d := &TreeNode{Val: 15}
	e := &TreeNode{Val: 7}

	a.Left = b
	a.Right = c
	c.Left = d
	c.Right = e

	t.Log(levelOrder(a))
}

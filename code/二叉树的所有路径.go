package code

import (
	"fmt"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var arr []int
	binaryTreeDfs(root, arr, &res)
	return res
}

func binaryTreeDfs(root *TreeNode, arr []int, res *[]string) {
	if root == nil {
		return
	}

	arr = append(arr, root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, formString(arr))
	}

	if root.Left != nil {
		binaryTreeDfs(root.Left, arr, res)
	}

	if root.Right != nil {
		binaryTreeDfs(root.Right, arr, res)
	}
}

func formString(arr []int) string {
	var res = make([]string, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		res = append(res, fmt.Sprintf("%d", arr[i]))
	}
	return strings.Join(res, "->")
}

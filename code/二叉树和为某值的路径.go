package code

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var arr = make([]int, 0)
	var res = make([][]int, 0)
	treeDfs(root, target, arr, &res)
	return res
}

func treeDfs(root *TreeNode, target int, arr []int, res *[][]int) {
	if root == nil {
		return
	}

	target -= root.Val
	arr = append(arr, root.Val)
	if root.Left == nil && root.Right == nil && 0 == target {
		// 需要对arr 进行拷贝
		*res = append(*res, append([]int(nil), arr...))
		return
	}
	treeDfs(root.Left, target, arr, res)
	treeDfs(root.Right, target, arr, res)
	arr = arr[:len(arr)-1]
}

package code

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var nodeList = []*TreeNode{root}
	for len(nodeList) != 0 {
		length := len(nodeList)
		var tmp = make([]int, length)
		for i, node := range nodeList {
			tmp[i] = node.Val
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
		}
		res = append(res, tmp)
		nodeList = nodeList[length:]
	}
	return res
}

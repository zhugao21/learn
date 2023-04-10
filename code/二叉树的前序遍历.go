package code

//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, root.Val)
		res = append(res, preorderTraversal(root.Left)...)
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}

func preorderTraversal01(root *TreeNode) []int {
	var res []int
	var nodeList []*TreeNode
	var cur = root

	for len(nodeList) != 0 || cur != nil {
		if cur != nil {
			res = append(res, cur.Val)
			nodeList = append(nodeList, cur.Right)
			cur = cur.Left
		} else {
			length := len(nodeList)
			node := nodeList[length-1]
			nodeList = nodeList[:length-1]
			cur = node
		}
	}
	return res
}

package code

//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

func inorderTraversal01(root *TreeNode) []int {
	var res []int
	var nodeList = []*TreeNode{}

	cur := root
	for len(nodeList) != 0 || cur != nil {
		if cur != nil {
			nodeList = append(nodeList, cur)
			cur = cur.Left
		} else {
			length := len(nodeList)
			cur = nodeList[length-1]
			nodeList = nodeList[:length-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

package code

//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, postorderTraversal(root.Left)...)
		res = append(res, postorderTraversal(root.Right)...)
		res = append(res, root.Val)
	}
	return res
}

// 后序遍历是 左右中。 可以看作是 中右左的反转
func postorderTraversal01(root *TreeNode) []int {
	var res []int
	var nodeList []*TreeNode
	var cur = root

	for cur != nil || len(nodeList) != 0 {
		if cur != nil {
			res = append(res, cur.Val)
			nodeList = append(nodeList, cur)
			cur = cur.Right
		} else {
			length := len(nodeList)
			node := nodeList[length-1]
			nodeList = nodeList[:length-1]
			cur = node.Left
		}
	}

	length := len(res)
	var nodes = make([]int, length)
	for i := 0; i < length; i++ {
		nodes[length-i-1] = res[i]
	}
	return nodes
}

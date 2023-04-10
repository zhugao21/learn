package code

///输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

func buildTree(preorder []int, inorder []int) *TreeNode {
	for i, val := range inorder {
		if val == preorder[0] {
			return &TreeNode{
				Val:   val,
				Left:  buildTree(preorder[1:i+1], inorder[0:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}

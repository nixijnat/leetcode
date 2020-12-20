package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	l := len(preorder)
	if l == 0 {
		return nil
	}
	// 先序遍历的第一个元素就是根节点
	root := &TreeNode{Val: preorder[0]}
	i := 0
	// 在中序中找到根节点
	for i = 0; i < l; i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 有左子树，就构造左子树
	if i != 0 {
		root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	}
	// 有右子树，就构造有子树
	if i != l-1 {
		root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	}
	return root
}

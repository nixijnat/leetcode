package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func upsideDownBinaryTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 新的头节点
	if root.Left == nil {
		return root
	}
	// 将左子树翻转
	nRoot := upsideDownBinaryTree1(root.Left)
	// 构建新子树
	root.Left.Left = root.Right
	root.Left.Right = root
	// 不要忘记将根节点置空
	root.Left, root.Right = nil, nil
	return nRoot
}

// 迭代
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 所有左子树入栈
	stk := make([]*TreeNode, 0, 8)
	for ; root != nil; root = root.Left {
		stk = append(stk, root)
	}
	var nRoot *TreeNode // 新的根节点
	var prev *TreeNode  // 上一个根节点
	for len(stk) > 0 {
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if nRoot == nil { // 新的根节点
			nRoot = root
			prev = root
			continue
		}
		// 翻转
		prev.Left = root.Right
		prev.Right = root
		// 置空当前节点
		root.Left, root.Right = nil, nil
		prev = root
	}
	return nRoot
}

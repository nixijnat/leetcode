package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1144
func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	flattenHelper(root)
}

// 返回尾节点
func flattenHelper(root *TreeNode) *TreeNode {
	tail := root
	if root.Left != nil { // 展开左子树
		right := root.Right
		root.Right = root.Left
		tail = flattenHelper(root.Left)
		tail.Right = right
		root.Left = nil // 很必要，将左儿子置空
	}
	if tail.Right != nil { // 展开右子树
		tail = flattenHelper(tail.Right)
	}
	return tail
}

// 深度展开版本
func flatten(root *TreeNode) {
	stack := make([]*TreeNode, 0, 8)
	var tail *TreeNode // 记录左子树展开后的尾节点
	for root != nil || len(stack) > 0 {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		right := root.Right
		if root.Left != nil {
			tail.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		// tail 为空，则表示叶子节点，则tail是root
		if tail == nil {
			tail = root
		}
		root = right // 展开右子树
		// 右子树不为空，则tail置空，表示存在新的 tail
		// 右子树为空，则tail 保持不变，还是左子树的 tail
		if root != nil {
			tail = nil
		}
	}
}

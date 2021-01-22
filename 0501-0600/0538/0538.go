package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 5389
func convertBST1(root *TreeNode) *TreeNode {
	find(root, 0)
	return root
}

// 返回 大值和
func find(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	root.Val += find(root.Right, sum)
	return find(root.Left, root.Val)
}

func convertBST(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0, 8)
	sum := 0
	res := root
	for root != nil || len(stack) != 0 {
		for ; root != nil; root = root.Right {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root.Val += sum
		sum = root.Val
		root = root.Left
	}
	return res
}

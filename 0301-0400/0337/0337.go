package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 3378
func rob(root *TreeNode) int {
	_, res := helper(root)
	return res
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lson, lcur := helper(root.Left)
	rson, rcur := helper(root.Right)
	return lcur + rcur, max(lson+rson+root.Val, lcur+rcur)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

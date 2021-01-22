package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res, _ := helper(root)
	return res
}

// 返回 子树中最大路径
// 返回 根结点单向连接的长度
func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lmax, lcur := helper(root.Left)
	rmax, rcur := helper(root.Right)
	return max(lcur+rcur, lmax, rmax), max(lcur, rcur) + 1
}

func max(as ...int) int {
	res := as[0]
	for _, v := range as {
		if v > res {
			res = v
		}
	}
	return res
}

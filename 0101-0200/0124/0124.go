package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max, _ := find(root)
	return max
}

// 返回值：子树内部的最大路径和 和
// 包含子树根节点相连的最大路径，为了与外部相连
func find(root *TreeNode) (int, int) {
	if root == nil {
		return -1000, -1000 // math.MinInt64 会溢出
	}
	lmax, lrootMax := find(root.Left)
	rmax, rrootMax := find(root.Right)
	// 包含子树根节点相连的最大路径，为了与外部相连
	rootMax := max(root.Val, lrootMax+root.Val, rrootMax+root.Val)
	// 子树内部的最大路径和
	mmax := max(lmax, rmax, rootMax, lrootMax+rrootMax+root.Val)
	return mmax, rootMax
}

func max(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}

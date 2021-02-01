package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	activeStk := make([]*TreeNode, 0, 8) // 活动栈
	bufStk := make([]*TreeNode, 0, 8)    // 缓存栈
	activeStk = append(activeStk, root)
	res := make([][]int, 0, 9)
	toRight := true // 方向
	for len(activeStk) != 0 {
		tmp := make([]int, 0, 8)
		// 遍历当前层级
		for len(activeStk) != 0 {
			node := activeStk[len(activeStk)-1]
			activeStk = activeStk[:len(activeStk)-1]
			tmp = append(tmp, node.Val)
			if toRight {
				// 从左往右，先压左节点，后压右节点
				if node.Left != nil {
					bufStk = append(bufStk, node.Left)
				}
				if node.Right != nil {
					bufStk = append(bufStk, node.Right)
				}
			} else {
				// 从右往左，先压右节点，后压左节点
				if node.Right != nil {
					bufStk = append(bufStk, node.Right)
				}
				if node.Left != nil {
					bufStk = append(bufStk, node.Left)
				}
			}
		}
		res = append(res, tmp)
		toRight = !toRight // 切换方向
		activeStk, bufStk = bufStk, activeStk
	}
	return res
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 4375
func pathSum(root *TreeNode, sum int) int {
	// 深度遍历
	nodeStack := make([]*TreeNode, 0, 8)
	flagStack := make([]bool, 0, 8) // 标记是否已经访问右节点
	res := 0
	total := 0
	prefix := make(map[int]int, 8) // 前缀和
	for root != nil || len(nodeStack) != 0 {
		// 深度搜索
		for ; root != nil; root = root.Left {
			prefix[total]++          // 前缀
			total += root.Val        // 加了当前节点就不是前缀和了
			res += prefix[total-sum] // 检查有多少满足结果
			nodeStack = append(nodeStack, root)
			flagStack = append(flagStack, true)
		}
		root = nodeStack[len(nodeStack)-1]
		// 是否访问右节点
		if !flagStack[len(flagStack)-1] {
			// 已经访问了左右节点，则pop
			total -= root.Val // 更新总和
			prefix[total]--   // 更新前缀和：去除
			nodeStack = nodeStack[:len(nodeStack)-1]
			flagStack = flagStack[:len(flagStack)-1]
			root = nil // 很重要
			continue
		}
		// 访问右节点
		flagStack[len(flagStack)-1] = false
		root = root.Right
	}
	return res
}

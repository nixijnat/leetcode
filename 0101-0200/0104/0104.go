package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth1(spaceIDs *TreeNode) int {
	return find(spaceIDs, 0)
}
func find(spaceIDs *TreeNode, cur int) int {
	if spaceIDs == nil {
		return cur
	}
	return max(find(spaceIDs.Left, cur+1), find(spaceIDs.Right, cur+1))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth2(root *TreeNode) int {
	// 结点栈
	stackNode := make([]*TreeNode, 0, 8)
	// 深度栈
	stackDepth := make([]int, 0, 8)
	max, depth := 0, 1
	for root != nil || len(stackNode) > 0 {
		for ; root != nil; root, depth = root.Left, depth+1 {
			stackNode = append(stackNode, root)
			stackDepth = append(stackDepth, depth)
		}
		root = stackNode[len(stackNode)-1]
		stackNode = stackNode[:len(stackNode)-1]
		depth = stackDepth[len(stackDepth)-1]
		stackDepth = stackDepth[:len(stackDepth)-1]
		fmt.Println(root.Val, depth, stackDepth)
		if depth > max {
			max = depth
		}
		root = root.Right
		depth++
	}
	return max
}
func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 结点队列
	queueNode := make([]*TreeNode, 0, 8)
	queueDepth := make([]int, 0, 8)
	// 深度队列
	queueNode = append(queueNode, root)
	queueDepth = append(queueDepth, 1)
	max := 0
	for len(queueNode) > 0 {
		cur := queueNode[0]
		queueNode = queueNode[1:]
		c := queueDepth[0]
		queueDepth = queueDepth[1:]
		if c > max {
			max = c
		}
		if cur.Left != nil {
			queueNode = append(queueNode, cur.Left)
			queueDepth = append(queueDepth, c+1)
		}
		if cur.Right != nil {
			queueNode = append(queueNode, cur.Right)
			queueDepth = append(queueDepth, c+1)
		}
	}
	return max
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 结点队列
	queueNode := make([]*TreeNode, 0, 8)
	queueNode = append(queueNode, root)
	depth := 0
	for len(queueNode) > 0 {
		l := len(queueNode)
		// 将该层次全部吐出
		for _, v := range queueNode {
			if v.Left != nil {
				queueNode = append(queueNode, v.Left)
			}
			if v.Right != nil {
				queueNode = append(queueNode, v.Right)
			}
		}
		// 去除上一层的元素
		queueNode = queueNode[l:]
		depth++
	}
	return depth
}

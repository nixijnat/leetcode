package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST1(root *TreeNode) bool {
	_, _, valid := isValid(root)
	return valid
}

func isValid(root *TreeNode) (int, int, bool) {
	min, max := root.Val, root.Val
	if root.Left != nil {
		// 判断左结点
		if root.Left.Val >= root.Val {
			return 0, 0, false
		}
		// 判断左子树
		tmin, tmax, valid := isValid(root.Left)
		if !valid || tmax >= root.Val {
			return 0, 0, false
		}
		min = tmin
	}
	if root.Right != nil {
		// 判断右结点
		if root.Right.Val <= root.Val {
			return 0, 0, false
		}
		// 判断右子树
		tmin, tmax, valid := isValid(root.Right)
		if !valid || tmin <= root.Val {
			return 0, 0, false
		}
		max = tmax
	}
	return min, max, true
}

func isValidBST_opt(root *TreeNode) bool {
	return isValid2(root, math.MinInt64, math.MaxInt64)
}

func isValid2(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValid2(root.Left, min, root.Val) && isValid2(root.Left, root.Val, max)
}

//
func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0, 8)
	last := &TreeNode{Val: math.MinInt64}
	for root != nil || len(stack) > 0 {
		// 一直向左走，走到该子树的起点
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		// 从栈中取出元素
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= last.Val {
			return false
		}
		last = root
		// 遍历右子树
		root = root.Right
	}
	return true
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal1(root *TreeNode) []int {
	res := []int{}
	helper(root, &res)
	return res
}
func helper(root *TreeNode, ns *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ns)
	helper(root.Right, ns)
	*ns = append((*ns), root.Val)
}

// 迭代：visit数组
func postorderTraversal_visit(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stk := make([]*TreeNode, 0, 8)
	stk = append(stk, root)
	visit := make([]bool, 1, 8) // 标识是否已经释放子节点
	for len(stk) > 0 {
		n := len(stk)
		root = stk[n-1]
		if visit[n-1] { //
			res = append(res, root.Val)
			stk = stk[:n-1]
			visit = visit[:n-1]
			continue
		}
		// 添加子节点
		visit[n-1] = true
		if root.Right != nil {
			stk = append(stk, root.Right)
			visit = append(visit, false)
		}
		if root.Left != nil {
			stk = append(stk, root.Left)
			visit = append(visit, false)
		}
	}
	return res
}

// 迭代：无 visit数组
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stk := make([]*TreeNode, 0, 8)
	var prev *TreeNode // 记录前一个访问的节点
	for root != nil || len(stk) > 0 {
		// 左儿子入栈
		for ; root != nil; root = root.Left {
			stk = append(stk, root)
		}
		n := len(stk)
		root = stk[n-1]
		// 没有右儿子则访问
		// 右儿子和前一个访问节点一样则访问：回溯
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			stk = stk[:n-1]
			prev = root
			root = nil // 置零
		} else {
			root = root.Right
		}
	}
	return res
}

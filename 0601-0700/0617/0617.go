package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 6178
func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// 如果有一颗树为空，直接返回非空树
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	// 合并子树
	t1.Val += t2.Val
	t1.Left = mergeTrees2(t1.Left, t2.Left)
	t1.Right = mergeTrees2(t1.Right, t2.Right)
	return t1
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	r, full := merge(t1, t2)
	if !full {
		return r
	}
	stack := make([]*TreeNode, 0, 8)
	stack = append(stack, t1, t2)
	for len(stack) != 0 {
		r1 := stack[len(stack)-2]
		r2 := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		r1.Left, full = merge(r1.Left, r2.Left)
		if full { // 如果子节点都不为空时 将子节点放入栈
			stack = append(stack, r1.Left, r2.Left)
		}
		r1.Right, full = merge(r1.Right, r2.Right)
		if full { // 如果子节点都不为空时 将子节点放入栈
			stack = append(stack, r1.Right, r2.Right)
		}
	}
	return t1
}

// 合并两个节点
// 返回合并后的节点
// 返回是否存在空节点
func merge(node1, node2 *TreeNode) (*TreeNode, bool) {
	if node1 == nil {
		return node2, false
	} else if node2 == nil {
		return node1, false
	}
	node1.Val += node2.Val
	return node1, true
}

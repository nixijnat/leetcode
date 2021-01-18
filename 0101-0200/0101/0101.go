package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric_recursion(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}
func isSymmetricHelper(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return isSymmetricHelper(l.Left, r.Right) && isSymmetricHelper(l.Right, r.Left)
}

func isSymmetric_queue(root *TreeNode) bool {
	q := make([]*TreeNode, 0, 8) // 队列，广度判断
	q = append(q, root, root)
	for len(q) > 0 {
		l, r := q[0], q[1]
		q = q[2:]
		if l == nil && r == nil {
			continue
		} else if l == nil || r == nil {
			return false
		} else if l.Val != r.Val {
			return false
		}
		q = append(q, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}
func isSymmetric(root *TreeNode) bool {
	stack := make([]*TreeNode, 0, 8) // 栈 深度判断
	stack = append(stack, root, root)
	for len(stack) > 0 {
		l, r := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if l == nil && r == nil {
			continue
		} else if l == nil || r == nil {
			return false
		} else if l.Val != r.Val {
			return false
		}
		stack = append(stack, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}

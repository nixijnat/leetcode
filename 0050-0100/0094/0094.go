package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 4)
	stack := make([]*TreeNode, 0, 4)
	cur := root
	for {
		if cur == nil {
			// 当前没有可判断的节点：左子树处理完毕
			// 从栈中取元素
			sl := len(stack)
			if sl == 0 {
				break
			}
			cur = stack[sl-1]
			stack = stack[:sl-1]
		} else {
			// 存在左子树则处理左子树
			if cur.Left != nil {
				stack = append(stack, cur)
				cur = cur.Left
				continue
			}
		}
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

func inorderTraversal_great(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 4)
	stack := make([]*TreeNode, 0, 4)
	for root != nil || len(stack) != 0 {
		for root != nil {
			// 将当前节点和左儿子放入栈中
			stack = append(stack, root)
			root = root.Left
		}
		// 从栈中取元素
		sl := len(stack)
		root = stack[sl-1]
		stack = stack[:sl-1]
		res = append(res, root.Val)
		// 遍历右子树
		root = root.Right
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 4)
	var inOrder func(bt *TreeNode)
	inOrder = func(bt *TreeNode) {
		if bt == nil {
			return
		}
		inOrder(bt.Left)          // 遍历左子树
		res = append(res, bt.Val) // 遍历根节点
		inOrder(bt.Right)         // 遍历右子树
	}
	inOrder(root)
	return res
}

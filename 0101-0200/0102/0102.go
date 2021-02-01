package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0, 8)
	queue := make([]*TreeNode, 0, 8)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]int, 0, length)
		// 遍历队列，不担心新加入元素
		for _, v := range queue {
			tmp = append(tmp, v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		res = append(res, tmp)
		queue = queue[length:]
	}
	return res
}

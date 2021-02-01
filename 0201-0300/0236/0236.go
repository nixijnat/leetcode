package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 分开找路径
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	l1 := find(root, p) // 分别找到两个节点的路径
	l2 := find(root, q)
	var res *TreeNode
	// 比较路径
	for i := 0; i < len(l1) && i < len(l2) && l1[i] == l2[i]; i++ {
		res = l1[i]
	}
	return res
}

func find(spaceRoot, a *TreeNode) []*TreeNode {
	stk := make([]*TreeNode, 0, 8)
	vis := make([]bool, 0, 8) // 记录是否已经访问右节点
	// 后续遍历，寻找路径
	for spaceRoot != nil || len(stk) != 0 {
		for ; spaceRoot != nil; spaceRoot = spaceRoot.Left {
			stk = append(stk, spaceRoot)
			vis = append(vis, false)
			if spaceRoot == a {
				return stk
			}
		}
		n := len(stk) - 1
		if vis[n] {
			stk = stk[:n]
			vis = vis[:n]
			continue
		}
		spaceRoot = stk[n].Right
		vis[n] = true
	}
	return stk
}

// 2365
func lowestCommonAncestor(root *TreeNode, p, q *TreeNode) *TreeNode {
	stknfd := make([]*TreeNode, 0, 8)
	stkfd := make([]*TreeNode, 0, 8)
	vis := make([]bool, 0, 8)
FIN:
	for root != nil || len(stknfd) != 0 {
		for ; root != nil; root = root.Left {
			stknfd = append(stknfd, root)
			vis = append(vis, false)
			// 找节点
			if root == p || root == q {
				if len(stkfd) == 0 { // 没有缓存则缓存
					stkfd = append(stkfd, stknfd...)
				} else { // 已经缓存则结束
					break FIN
				}
			}
		}
		n := len(stknfd) - 1
		if vis[n] {
			stknfd = stknfd[:n]
			vis = vis[:n]
			continue
		}
		root = stknfd[n].Right
		vis[n] = true
	}
	// 比较路径
	var res *TreeNode
	for i := 0; i < len(stkfd) && i < len(stknfd) &&
		stkfd[i] == stknfd[i]; i++ {
		res = stkfd[i]
	}
	return res
}

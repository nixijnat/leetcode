package main

import (
	"container/list"
	"reflect"
	"testing"
)

func makeTree(arr []int) *TreeNode {
	l := len(arr)
	if l == 0 {
		return nil
	}
	var mk func(i int) *TreeNode
	mk = func(i int) *TreeNode {
		if i >= l || arr[i] == 0 {
			return nil
		}
		root := &TreeNode{arr[i], nil, nil}
		root.Left = mk(2*i + 1)
		root.Right = mk(2*i + 2)
		return root
	}
	return mk(0)
}
func resovleTree(t *TreeNode) [][]int {
	l := list.New()
	l.PushBack(t)
	res := make([][]int, 0, 4)
	for l.Len() != 0 {
		v := l.Remove(l.Front()).(*TreeNode)
		tmp := []int{v.Val, -1, -1}
		if v.Left != nil {
			l.PushBack(v.Left)
			tmp[1] = v.Left.Val
		}
		if v.Right != nil {
			l.PushBack(v.Right)
			tmp[2] = v.Right.Val
		}
		res = append(res, tmp)
	}
	return res
}

func TestFunc(t *testing.T) {
	tr := makeTree([]int{3, 5, 1, 6, 2, 9, 8, 0, 0, 7, 4})
	for _, c := range []struct {
		input1 int
		input2 int
		expect int
	}{
		{5, 1, 3},
		{5, 4, 5},
		{6, 4, 5},
		{7, 8, 3},
	} {
		p := findNode(tr, c.input1)
		q := findNode(tr, c.input2)
		v := lowestCommonAncestor(tr, p, q)
		ex := findNode(tr, c.expect)
		t.Log(c, v, reflect.DeepEqual(v, ex))
	}
}

func findNode(root *TreeNode, a int) *TreeNode {
	stk := make([]*TreeNode, 0, 8)
	for root != nil || len(stk) != 0 {
		for ; root != nil; root = root.Left {
			if root.Val == a {
				return root
			}
			stk = append(stk, root)
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		root = root.Right
	}
	return nil
}

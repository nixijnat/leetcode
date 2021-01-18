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
	for _, c := range []struct {
		input1 []int
		expect bool
	}{
		{[]int{1, 2, 2, 3, 4, 4, 3}, true},
		{[]int{1, 2, 2, 0, 4, 0, 4}, false},
		{[]int{1}, true},
		{[]int{1, 2, 0}, false},
		{[]int{1, 0, 2}, false},
	} {
		v := isSymmetric(makeTree(c.input1))
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

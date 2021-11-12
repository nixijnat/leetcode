package q35

import (
	"leetcode/common/cmp"
	"testing"
)

func makeList(is [][]int) *Node {
	if len(is) == 0 {
		return nil
	}
	dummy := &Node{}
	n := dummy
	ns := []*Node{}
	for _, i := range is {
		n.Next = &Node{Val: i[0]}
		ns = append(ns, n.Next)
		n = n.Next
	}
	n = dummy
	for _, i := range is {
		if i[1] != -1 {
			n.Next.Random = ns[i[1]]
		}
		n = n.Next
	}
	return dummy.Next
}

func resovleList(n *Node) [][]int {
	if n == nil {
		return nil
	}
	ns := make(map[*Node]int, 0)
	for m, i := n, 0; m != nil; m, i = m.Next, i+1 {
		ns[m] = i
	}
	res := [][]int{}
	for m := n; m != nil; m = m.Next {
		if m.Random == nil {
			res = append(res, []int{m.Val, -1})
		} else {
			res = append(res, []int{m.Val, ns[m.Random]})
		}
	}
	return res
}

func TestFunc(t *testing.T) {
	for _, cas := range []struct {
		input [][]int
	}{
		{[][]int{}},
		{[][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
		{[][]int{{1, 1}, {2, 1}}},
		{[][]int{{3, -1}, {3, 0}, {3, -1}}},
	} {
		n := makeList(cas.input)
		t.Log(resovleList(n))
		res := resovleList(copyRandomList(n))
		if !cmp.IsEqualIntSlice2D(cas.input, res) {
			t.Fatalf("input: %v, but get: %v", cas.input, res)
		}
	}
}

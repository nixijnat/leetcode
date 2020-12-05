package main

import (
	"fmt"
	"math"
	"testing"
)

func genList(v int) *ListNode {
	if v == 0 {
		return new(ListNode)
	}
	var ret, cur *ListNode
	tmp := 0
	for v != 0 {
		tmp, v = v%10, v/10
		node := &ListNode{
			Val: tmp,
		}
		if cur != nil {
			cur.Next = node
		} else {
			ret = node
		}
		cur = node
	}
	return ret
}

func printNode(v *ListNode) int {
	sum := 0
	for i := 0; v != nil; v = v.Next {
		sum = sum + v.Val*int(math.Pow10(i))
		i++
	}
	return sum
}

func TestTwoSum(t *testing.T) {
	for _, c := range []struct {
		a int
		b int
	}{
		{0, 0},
		{0, 5},
		{0, 10},
		{4, 13},
		{7, 18},
		{47, 345},
		{2222, 18},
		{9999999, 9999},
	} {
		na := genList(c.a)
		nb := genList(c.b)
		fmt.Printf("%d + %d = %d\n", printNode(na), printNode(nb), printNode(addTwoNumbers(na, nb)))
	}
}

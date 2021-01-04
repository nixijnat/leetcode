package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	v := Constructor()
	v.Push(-2)
	v.Push(0)
	v.Push(-3)
	t.Log(v.GetMin())
	v.Pop()
	t.Log(v.Top())
	t.Log(v.GetMin())
}

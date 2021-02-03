package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return this.rserialize(root, "")
}

func (this *Codec) rserialize(root *TreeNode, str string) string {
	if root == nil {
		str += ","
		return str
	}
	// 先序遍历
	str += strconv.Itoa(root.Val) + ","
	str = this.rserialize(root.Left, str)
	str = this.rserialize(root.Right, str)
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.l = strings.Split(data, ",")
	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	if this.l[0] == "" {
		this.l = this.l[1:] // 移除字符串
		return nil
	}
	// 先序遍历
	num, _ := strconv.Atoi(this.l[0])
	node := &TreeNode{Val: num}
	this.l = this.l[1:] // 移除字符串
	node.Left = this.rdeserialize()
	node.Right = this.rdeserialize()
	return node
}

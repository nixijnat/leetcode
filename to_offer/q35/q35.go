package q35

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	dummySrc := &Node{Next: head}
	dummyDst := &Node{}
	ns := make(map[*Node]*Node)
	for n, m := dummySrc, dummyDst; n.Next != nil; n, m = n.Next, m.Next {
		m.Next = &Node{Val: n.Next.Val, Random: n.Next.Random}
		ns[n.Next] = m.Next
	}
	for i, n := 0, dummyDst.Next; n != nil; i, n = i+1, n.Next {
		n.Random = ns[n.Random]
	}
	return dummyDst.Next
}

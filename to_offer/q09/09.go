package q09

// 09.用两个栈实现队列

type stack []int

func (s *stack) push(value int) {
	*s = append(*s, value)
}

func (s *stack) pop() int {
	olen := s.length() - 1
	res := (*s)[olen]
	*s = (*s)[:olen]
	return res
}

func (s *stack) length() int {
	return len(*s)
}

type CQueue struct {
	in  stack
	out stack
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in.push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.out.length() == 0 {
		for this.in.length() != 0 {
			this.out.push(this.in.pop())
		}
	}
	if this.out.length() == 0 {
		return -1
	}
	return this.out.pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

package q30

// 30. 包含min函数的栈

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

func (s *stack) top() int {
	return (*s)[s.length()-1]
}

type MinStack struct {
	raw stack
	min stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Push(x int) {
	this.raw.push(x)
	if this.min.length() == 0 {
		this.min.push(x)
	} else {
		this.min.push(min(x, this.min.top()))
	}
}

func (this *MinStack) Pop() {
	this.raw.pop()
	this.min.pop()
}

func (this *MinStack) Top() int {
	return this.raw.top()
}

func (this *MinStack) Min() int {
	return this.min.top()
}

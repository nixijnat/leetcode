package main

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	ml := len(this.min)
	if ml == 0 || this.min[ml-1] > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[ml-1])
	}
}

func (this *MinStack) Pop() {
	l := len(this.data)
	this.data = this.data[:l-1]
	this.min = this.min[:l-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

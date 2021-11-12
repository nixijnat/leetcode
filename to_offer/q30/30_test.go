package q30

import "testing"

func TestFunc(t *testing.T) {
	c := Constructor()
	c.Push(-2)
	c.Push(0)
	c.Push(-3)
	if res := c.Min(); res != -3 {
		t.Fatalf("expect -3, get %d", res)
	}
	c.Pop()
	if res := c.Top(); res != 0 {
		t.Fatalf("expect 0, get %d", res)
	}
	if res := c.Min(); res != -2 {
		t.Fatalf("expect -2, get %d", res)
	}
}

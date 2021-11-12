package q09

import (
	"testing"
)

func TestFunc(t *testing.T) {
	c := Constructor()
	c.AppendTail(3)
	if res := c.DeleteHead(); res != 3 {
		t.Fatalf("expect 3, get %d", res)
	}
	if res := c.DeleteHead(); res != -1 {
		t.Fatalf("expect -1, get %d", res)
	}
}

func TestFunc2(t *testing.T) {
	c := Constructor()
	if res := c.DeleteHead(); res != -1 {
		t.Fatalf("expect -1, get %d", res)
	}
	c.AppendTail(5)
	c.AppendTail(2)
	if res := c.DeleteHead(); res != 5 {
		t.Fatalf("expect 5, get %d", res)
	}
	if res := c.DeleteHead(); res != 2 {
		t.Fatalf("expect 2, get %d", res)
	}
}

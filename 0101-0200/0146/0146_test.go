package main

import (
	"fmt"
	"testing"
)

func TestFunc2(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		c := Constructor(2)
		c.Put(1, 1)
		c.Put(2, 2)
		fmt.Println(1, c.Get(1), 1)
		c.Put(3, 3)
		fmt.Println(2, c.Get(2), -1)
		c.Put(4, 4)
		fmt.Println(1, c.Get(1), -1)
		fmt.Println(3, c.Get(3), 3)
		fmt.Println(4, c.Get(4), 4)
	})
	t.Run("2", func(t *testing.T) {
		c := Constructor(2)
		c.Put(1, 1)
		c.Put(2, 2)
		fmt.Println(1, c.Get(1), 1)
		c.Put(2, 22)
		c.Put(3, 3)
		fmt.Println(2, c.Get(2), 22)
		c.Put(4, 4)
		fmt.Println(1, c.Get(1), -1)
		fmt.Println(2, c.Get(2), 22)
		fmt.Println(3, c.Get(3), -1)
		fmt.Println(4, c.Get(4), 4)
	})
}

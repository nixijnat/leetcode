package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []string
		expect [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{[]string{"ate", "eat", "tea"},
				[]string{"nat", "tan"},
				[]string{"bat"}}},
	} {
		v := groupAnagrams(c.input1)
		t.Log(c, v)
	}
}

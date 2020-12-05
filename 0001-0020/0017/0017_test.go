package main

import "testing"

func TestThreeSum(t *testing.T) {
	for _, cas := range []struct {
		input string
	}{
		{"23"},
		{"234"},
	} {
		v := letterCombinations(cas.input)
		t.Log(cas, v)
	}
}

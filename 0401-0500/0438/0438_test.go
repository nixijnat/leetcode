package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	longStr := strings.Repeat("a", 10000)
	for _, c := range []struct {
		input1 string
		input2 string
		expect []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"", "", []int{}},
		{"abacbabc", "abc", []int{1, 2, 3, 5}},
		{longStr + "b" + longStr, longStr, []int{0, 10001}},
	} {
		v := findAnagrams(c.input1, c.input2)
		t.Log(v, reflect.DeepEqual(c.expect, v))
	}
}

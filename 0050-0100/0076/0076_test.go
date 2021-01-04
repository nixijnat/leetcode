package main

import (
	"reflect"
	"testing"
)

func TestFunc_nocount(t *testing.T) {
	for _, c := range []struct {
		input1 string
		input2 string
		expect string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"ADOBECODEBANC", "ABCQ", ""},
		{"A", "A", "A"},
		{"A", "AA", "A"},
		{"", "", ""},
	} {
		v := minWindow_nocount(c.input1, c.input2)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}
func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		input2 string
		expect string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"ADOBECODEBANC", "ABCQ", ""},
		{"A", "A", "A"},
		{"A", "AA", ""},
		{"", "", ""},
	} {
		v := minWindow(c.input1, c.input2)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

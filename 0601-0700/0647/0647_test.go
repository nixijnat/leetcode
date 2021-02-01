package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		expect int
	}{
		{"aaa", 6},
		{"abc", 3},
	} {
		v := countSubstrings(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

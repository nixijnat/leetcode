package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		expect []string
	}{
		{"()())()", []string{"()()()", "(())()"}},
		{"(a)())()", []string{"(a)()()", "(a())()"}},
		{")(", []string{""}},
	} {
		v := removeInvalidParentheses(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

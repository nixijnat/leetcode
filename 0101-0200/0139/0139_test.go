package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		input2 []string
		expect bool
	}{
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "and", "sand", "dog", "cat"}, false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, false},
	} {
		v := wordBreak(c.input1, c.input2)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

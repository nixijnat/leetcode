package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		input2 []string
		expect []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{}},
		{"aaa", []string{"a", "a"}, []int{1, 0}},
		{"dddddddddddd", []string{"dddd", "dddd"}, []int{1, 0}},
	} {
		v := findSubstring(c.input1, c.input2)
		t.Log(c, v)
	}
}

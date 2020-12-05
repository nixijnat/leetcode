package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []string
		expect string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", "", "car"}, ""},
		{[]string{"a", "ab"}, "a"},
	} {
		v := longestCommonPrefix(c.input)
		t.Log(c, v, v == c.expect)
	}
}

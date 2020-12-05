package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		expect string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
	} {
		v := countAndSay(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

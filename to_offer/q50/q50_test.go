package q50

import "testing"

func TestFunc(t *testing.T) {
	for _, cas := range []struct {
		str    string
		expect byte
	}{
		{"", ' '},
		{"z", 'z'},
		{"abaccdeff", 'b'},
		{"leetcode", 'l'},
	} {
		res := firstUniqChar(cas.str)
		if res != cas.expect {
			t.Fatalf("input: %v, expect: %c, but get: %c", cas.str, cas.expect, res)
		}
	}
}

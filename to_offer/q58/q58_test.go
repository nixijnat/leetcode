package q58

import "testing"

func TestFunc(t *testing.T) {
	for _, cas := range []struct {
		str    string
		pos    int
		expect string
	}{
		{"", 0, ""},
		{"abcdefg", -1, "abcdefg"},
		{"abcdefg", 0, "abcdefg"},
		{"abcdefg", 2, "cdefgab"},
		{"abcdefg", 4, "efgabcd"},
		{"abcdefg", 7, "abcdefg"},
		{"abcdefg", 8, "abcdefg"},
		{"lrloseumgh", 6, "umghlrlose"},
	} {
		res := reverseLeftWords(cas.str, cas.pos)
		if res != cas.expect {
			t.Fatalf("input: %v,%d, expect: %v, but get: %v", cas.str, cas.pos, cas.expect, res)
		}
	}
}

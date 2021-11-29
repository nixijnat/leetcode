package q05

import "testing"

func TestFunc(t *testing.T) {
	for _, cas := range []struct {
		input  string
		expect string
	}{
		{"", ""},
		{" ", "%20"},
		{"  ", "%20%20"},
		{"We are happy.", "We%20are%20happy."},
	} {
		res := replaceSpace(cas.input)
		if res != cas.expect {
			t.Fatalf("input: %v, expect: %v, but get: %v", cas.input, cas.expect, res)
		}
	}
}

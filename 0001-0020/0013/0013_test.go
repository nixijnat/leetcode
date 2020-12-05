package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"VI", 6},
		{"VII", 7},
		{"VIII", 8},
		{"IX", 9},
		{"X", 10},
		{"XI", 11},
		{"XII", 12},
		{"XIII", 13},
		{"XIV", 14},
		{"XV", 15},
		{"XVI", 16},
		{"XVII", 17},
		{"XVIII", 18},
		{"XIX", 19},
		{"XX", 20},
		{"XXX", 30},
		{"XL", 40},
		{"L", 50},
		{"LX", 60},
		{"LXX", 70},
		{"LXXX", 80},
		{"XC", 90},
		{"XCIX", 99},
		{"C", 100},
		{"CI", 101},
		{"CII", 102},
		{"CXCIX", 199},
		{"CC", 200},
		{"CCC", 300},
		{"CD", 400},
		{"D", 500},
		{"DC", 600},
		{"DCCC", 800},
		{"CM", 900},
		{"M", 1000},
		{"MCD", 1400},
		{"MCDXXXVII", 1437},
		{"MD", 1500},
		{"MDCCC", 1800},
		{"MDCCCLXXX", 1880},
		{"MCM", 1900},
		{"MM", 2000},
		{"MMM", 3000},
		{"MMMCCCXXXIII", 3333},
		{"MMMM", 4000},
	} {
		v := romanToInt(c.input)
		t.Log(c, v, v == c.expect)
	}
}

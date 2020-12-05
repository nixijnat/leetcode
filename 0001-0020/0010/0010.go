package main

import (
	"fmt"
	"strings"
)

func isMatch(s string, p string) bool {
	switch {
	case s == "":
		return true
	case p == "":
		return false
	}
	if strings.Contains(p, ".*") {
		return true
	}
	i, j := 0, 0
	pStart := 0
	for j = pStart; i < len(s) && j < len(p); {
		if s == "abcabccce" {
			fmt.Printf("%d, %d, %c, %c\n", i, j, s[i], p[j])
		}

		if s[i] == p[j] || p[j] == '.' {
			i++
			j++
			continue
		}
		if i == 0 {
			j++
			continue
		}
		if p[j] != '*' {
			i = 0
			pStart++
			j = pStart
			continue
		}

		for ; i > 0 && i < len(s); i++ {
			if s[i-1] != s[i] {
				break
			}
		}
		if i == len(s) {
			return true
		}
		j++
	}
	return i == len(s)
}

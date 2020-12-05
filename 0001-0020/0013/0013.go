package main

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	length := len(s)
	ret := 0
	for i := 0; i < length; i++ {
		inc := 0
		switch s[i] {
		case 'I':
			if i+1 < length {
				if s[i+1] == 'V' {
					inc = 4
				} else if s[i+1] == 'X' {
					inc = 9
				}
			}
			if inc == 0 {
				ret += 1
			}
		case 'V':
			ret += 5
		case 'X':
			if i+1 < length {
				if s[i+1] == 'L' {
					inc = 40
				} else if s[i+1] == 'C' {
					inc = 90
				}
			}
			if inc == 0 {
				ret += 10
			}
		case 'L':
			ret += 50
		case 'C':
			if i+1 < length {
				if s[i+1] == 'D' {
					inc = 400
				} else if s[i+1] == 'M' {
					inc = 900
				}
			}
			if inc == 0 {
				ret += 100
			}
		case 'D':
			ret += 500
		case 'M':
			ret += 1000
		default:
			return 0
		}
		if inc != 0 {
			ret += inc
			i++
		}
	}
	return ret
}

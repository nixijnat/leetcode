package cmp

func IsEqualIntSlice(a []int, b []int) bool {
	al, bl := len(a), len(b)
	if al != bl {
		return false
	}
	for i, ae := range a {
		if ae != b[i] {
			return false
		}
	}
	return true
}

func IsEqualIntSlice2D(a [][]int, b [][]int) bool {
	al, bl := len(a), len(b)
	if al != bl {
		return false
	}
	for i, ae := range a {
		if !IsEqualIntSlice(ae, b[i]) {
			return false
		}
	}
	return true
}

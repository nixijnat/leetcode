package main

func hammingDistance_direct(x int, y int) int {
	res := 0
	for ; x != 0 || y != 0; x, y = x>>1, y>>1 {
		if x&0x1 != y&0x1 {
			res++
		}
	}
	return res
}

func hammingDistance(x int, y int) int {
	res := 0
	for i := x ^ y; i != 0; i &= i - 1 {
		res++
	}
	return res
}

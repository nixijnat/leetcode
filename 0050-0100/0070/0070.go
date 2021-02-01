package main

// 7023
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	last1, last2 := 1, 2
	for i := 3; i <= n; i++ {
		last1, last2 = last2, last1+last2
	}
	return last2
}

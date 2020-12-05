package main

func maxArea(height []int) int {
	var ret, tmp int
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			tmp = height[i] * (j - i)
			i++
		} else {
			tmp = height[j] * (j - i)
			j--
		}
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}

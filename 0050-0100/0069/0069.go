package main

func mySqrt1(x int) int {
	if x <= 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	count := 0
	for i := x; i != 1; count++ {
		i >>= 1
	}
	// 找到这个初始化
	start := 1
	for count = count >> 1; count > 0; count-- {
		start <<= 1
	}
	// 向上找
	for {
		tmp := start * start
		if tmp == x {
			return start
		} else if tmp > x {
			return start - 1
		}
		start++
	}
}
func mySqrt2(x int) int {
	if x <= 0 {
		return 0
	} else if x < 4 {
		return 1
	}
	mid := 0
	for low, high := 2, x; low <= high; {
		mid = low + (high-low)>>1
		// 当前数
		tmp1 := mid * mid
		if tmp1 == x {
			break
		}
		// 下一个数
		tmp2 := (mid + 1) * (mid + 1)
		if tmp2 == x {
			mid = mid + 1
			break
		}
		if tmp1 < x && tmp2 > x {
			break
		} else if tmp1 > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return mid
}

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	} else if x < 4 {
		return 1
	}
	res := 0
	for low, high := 2, x; low <= high; {
		mid := low + (high-low)>>1
		if mid*mid <= x {
			res = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}

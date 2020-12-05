package main

func reverse(x int) int {
	// 个位直接返回
	if x >= -9 && x <= 9 {
		return x
	}
	const MAX = 1<<31 - 1 // 极值
	ret, cal := int64(0), int64(x)
	for ; cal != 0; cal /= 10 {
		ret = ret*10 + cal%10
		if ret > MAX || ret < -MAX {
			return 0
		}
	}
	return int(ret)
}

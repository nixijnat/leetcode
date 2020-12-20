package main

func subsets1(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	l = 1 << l // 2 ^ n
	res := make([][]int, 0, l)
	res = append(res, []int{})
	for _, v1 := range nums {
		// 复制旧的，
		// 然后再加上 v1
		for _, v2 := range res {
			tmp := make([]int, len(v2), len(v2)+1)
			copy(tmp, v2)
			tmp = append(tmp, v1)
			res = append(res, tmp)
		}
	}
	return res
}

func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	n = 1 << n
	res := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, 0, 2)
		for j, v := range nums {
			// 1 << j & i 判断 j 位置上数字存不存在
			if 1<<j&i != 0 {
				tmp = append(tmp, v)
			}
		}
		res = append(res, tmp)
	}
	return res
}

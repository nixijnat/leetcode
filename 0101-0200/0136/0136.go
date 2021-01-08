package main

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	for val, count := range m {
		if count == 1 {
			return val
		}
	}
	return 0
}

func singleNumber3(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}

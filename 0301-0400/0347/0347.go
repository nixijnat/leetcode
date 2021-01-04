package main

func topKFrequent(nums []int, k int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	counts := make(map[int]int, length)
	for _, v := range nums {
		counts[v]++
	}
	values := make([][2]int, 0, length)
	for val, count := range counts {
		values = append(values, [2]int{val, count})
	}
	// 找到第k个元素
	for low, high := 0, len(values)-1; low < high; {
		mid := partition(values, low, high)
		switch {
		case mid == k-1:
			break
		case mid < k-1:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, values[i][0])
	}
	return res
}
func partition(v [][2]int, low, high int) int {
	vt := v[low]
	for low < high {
		for ; low < high && v[high][1] <= vt[1]; high-- {
		}
		v[low] = v[high]
		for ; low < high && v[low][1] >= vt[1]; low++ {
		}
		v[high] = v[low]
	}
	v[low] = vt
	return low
}

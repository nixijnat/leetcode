package main

func sortColors_insert(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	// insert 不符合题意的原地操作
	res := make([]int, 1, l)
	res[0] = nums[0]
	for i := 1; i < l; i++ {
		j := 0
		for ; j < len(res) && nums[i] >= res[j]; j++ {
		}
		if j == len(res) {
			res = append(res, nums[i])
		} else {
			res = append(res[:j+1], res[j:]...)
			res[j] = nums[i]
		}
	}
	return res
}

func sortColors_select(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	// select
	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return
}

func sortColors_buble(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	// buble
	for i := 0; i < l-1; i++ {
		for j := l - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return
}

func sortColors_quick(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	// quick
	mid := partition(nums)
	if mid >= 1 {
		sortColors(nums[0:mid])
	}
	if mid <= l-2 {
		sortColors(nums[mid+1 : l])
	}
	return
}

func partition(ar []int) int {
	pivot := ar[0]
	low, high := 0, len(ar)-1
	for low < high {
		// 先高往低
		// 要用 >= , 否则死循环
		for ; low < high && ar[high] >= pivot; high-- {
		}
		ar[low] = ar[high]
		for ; low < high && ar[low] <= pivot; low++ {
		}
		ar[high] = ar[low]
	}
	ar[low] = pivot
	return low
}

func sortColors_count(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	var counts [3]int
	for _, v := range nums {
		counts[v]++
	}
	tmp := nums[:0]
	for v, c := range counts {
		for i := 0; i < c; i++ {
			tmp = append(tmp, v)
		}
	}
}

func sortColors1(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	p0, p2 := 0, l-1
	// i 遇到 p2 则停止
	for i := 0; i <= p2; i++ {
		// 循环处理，p2 与 i 的2 2 交换情况
		for i <= p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		// 正常情况 或者 p2 把 0交换到 i
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

func sortColors(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	p0, p1 := 0, 0
	// i 遇到 p2 则停止
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			// p0 < p1 必然会把 1 交换到 i，所以应该交换回去
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			// 不能出现 p0 > p1
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

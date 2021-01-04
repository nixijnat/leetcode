package main

import "sort"

func majorityElement1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	sort.Ints(nums)
	return nums[(length-1)/2]
}

func majorityElement2(nums []int) int {
	n := len(nums)
	mp := make(map[int]int, n)
	for _, v := range nums {
		mp[v]++
	}
	n = n / 2
	for k, cnt := range mp {
		if cnt > n {
			return k
		}
	}
	return 0
}

func majorityElement3(nums []int) int {
	n := len(nums)
	k := (n - 1) / 2
	for low, high := 0, n-1; low <= high; {
		mid := partition(nums, low, high)
		switch {
		case mid == k:
			return nums[mid]
		case mid < k:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return 0
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]
	for low < high {
		for ; low < high && nums[high] >= pivot; high-- {
		}
		nums[low] = nums[high]
		for ; low < high && nums[low] <= pivot; low++ {
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}

func majorityElement(nums []int) int {
	candiate, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			candiate = v
		}
		if v == candiate {
			count++
		} else {
			count--
		}
	}
	return candiate
}

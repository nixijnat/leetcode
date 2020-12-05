package main

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	low, high := 0, l-1
	for low <= high {
		mid := (high + low) >> 1
		switch {
		case target == nums[mid]:
			return mid
		case target < nums[mid]:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	// 没找到返回 low
	return low
}

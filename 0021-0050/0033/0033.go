package main

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	low, high := 0, l-1
	// 注意 >= 和 <=
	for low <= high {
		mid := (high + low) >> 1
		lowPart := false
		switch {
		case target == nums[mid]: // 找到了
			return mid
		case nums[mid] >= nums[low]: // mid >= low
			// mid 在 顺序区
			lowPart = target >= nums[low] && target < nums[mid]
		default: // mid < low
			// mid 在 乱序区
			lowPart = !(target > nums[mid] && target <= nums[high])
		}
		// 找是不是在低位即可
		if lowPart {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

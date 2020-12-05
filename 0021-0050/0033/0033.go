package main

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	low, high := 0, l-1
	for low < high {
		mid := (high + low) >> 1
		if target == nums[mid] { // 找到了
			return mid
		}
		// 找是不是在低位即可
		lowPart := true
		if nums[mid] >= nums[low] {
			if nums[low] == target {
				return low
			}
			// mid 在 顺序区
			lowPart = target > nums[low] && target < nums[mid]
		} else {
			if nums[high] == target {
				return high
			}
			// mid 在 乱序区
			lowPart = !(target > nums[mid] && target < nums[high])
		}
		if lowPart {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	// 单独把 low==high 拎出来，避免在循环中多余的处理
	if low == high && nums[low] == target {
		return low
	}
	return -1
}

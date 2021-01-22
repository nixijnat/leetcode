package main

// 81
func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		lowPart := true
		switch {
		case nums[mid] == target:
			return true
		case nums[mid] == nums[low]: // 缩小访问
			low++
			continue
		case nums[mid] > nums[low]:
			lowPart = nums[mid] > target && target >= nums[low]
		default:
			lowPart = !(nums[high] >= target && target > nums[mid])
		}
		if lowPart {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

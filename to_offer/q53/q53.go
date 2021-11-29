package q53

// 53 - I. 在排序数组中查找数字 I

func search(nums []int, target int) int {
	left := -1
	for i, j := 0, len(nums)-1; i <= j; {
		mid := i + (j-i)/2
		if nums[mid] >= target {
			j = mid - 1
			left = mid
		} else {
			i = mid + 1
		}
	}
	if left == -1 {
		return 0
	}
	right := -1
	for i, j := 0, len(nums)-1; i <= j; {
		mid := i + (j-i)/2
		if nums[mid] <= target {
			i = mid + 1
			right = mid
		} else {
			j = mid - 1
		}
	}
	res := right - left + 1
	if res < 0 {
		return 0
	}
	return res
}

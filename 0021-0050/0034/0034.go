package main

func searchRange1(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	} else if l == 1 { // 一个元素
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	flag := -1
	// 找目标值
	low, high := 0, l-1
	for low <= high {
		mid := (high + low) >> 1
		if target == nums[mid] {
			flag = mid
			break
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if flag == -1 { // 不存在这个值
		return []int{-1, -1}
	}
	ret := []int{flag, flag}
	// 如果 flag 的上一个数也是target，那就继续向上找结尾
	if flag > 0 && nums[flag-1] == target {
		for low, high = 0, flag-1; low <= high; {
			mid := (low + high) >> 1
			if target == nums[mid] {
				ret[0] = mid
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	// 如果 flag 的下一个数也是target，那就继续向下找结尾
	if flag < l-1 && nums[flag+1] == target {
		for low, high = flag+1, l-1; low <= high; {
			mid := (low + high) >> 1
			if target == nums[mid] {
				ret[1] = mid
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return ret
}

func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	} else if l == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	ret := []int{-1, -1}
	for low, high := 0, l-1; low <= high; {
		mid := (low + high) >> 1
		if target < nums[mid] {
			high = mid - 1
		} else if target == nums[mid] {
			ret[0] = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	for low, high := 0, l-1; low <= high; {
		mid := (low + high) >> 1
		if target > nums[mid] {
			low = mid + 1
		} else if target == nums[mid] {
			ret[1] = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ret
}

package main

func findUnsortedSubarray1(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	// max 存储极大数索引 / min 存储极小数索引
	max, min := -1, -1
	// 从前往后，遇见向下趋势，则前一个元素为极大值
	for i := 1; i < length; i++ {
		// 向下趋势
		if nums[i] < nums[i-1] {
			if max == -1 || nums[i-1] > nums[max] {
				max = i - 1 // 更新极大值索引
			}
		}
	}
	// 没有向下趋势
	if max == -1 {
		return 0
	}
	// 从后往前，遇见向上趋势，则后一个元素为极小值
	for i := length - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			if min == -1 || nums[i+1] < nums[min] {
				min = i + 1
			}
		}
	}
	// 从前往后找第一个比极小值大的数
	for i := 0; i < length; i++ {
		if nums[i] > nums[min] {
			min = i
			break
		}
	}
	// 从后往前找第一个比极大值小的数
	for i := length - 1; i >= 0; i-- {
		if nums[i] < nums[max] {
			max = i
			break
		}
	}
	return max - min + 1
}
func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	// min 从后往前 所以是 length-1
	min, max := nums[length-1], nums[0]
	start, end := 0, -1
	for i := 0; i < length; i++ {
		// 从前往后记录最大值
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
		// 从后往前记录最小值
		if nums[length-i-1] <= min {
			min = nums[length-i-1]
		} else {
			start = length - i - 1
		}
	}
	return end - start + 1
}

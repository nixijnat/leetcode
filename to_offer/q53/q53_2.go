package q53

// 53 - II. 0～n-1中缺失的数字
func missingNumberN(nums []int) int {
	var i int
	for i = 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return i
}

func missingNumber(nums []int) int {
	var i int
	for j := len(nums) - 1; i <= j; {
		mid := i + (j-i)/2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

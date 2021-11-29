package q04

// 04. 二维数组中的查找

func binarySearch(nums []int, target int) bool {
	for i, j := 0, len(nums)-1; i <= j; {
		mid := i + (j-i)>>1
		switch {
		case nums[mid] == target:
			return true
		case nums[mid] < target:
			i = mid + 1
		default:
			j = mid - 1
		}
	}
	return false
}

func findNumberIn2DArrayBinarySearch(matrix [][]int, target int) bool {
	for i := range matrix {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

// linear search : 线性查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	for i, j := m-1, 0; i >= 0 && j < n; {
		switch {
		case matrix[i][j] == target:
			return true
		case matrix[i][j] < target:
			j++
		default:
			i--
		}
	}
	return false
}

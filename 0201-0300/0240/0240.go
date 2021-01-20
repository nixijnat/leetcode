package main

func searchMatrix1(matrix [][]int, target int) bool {
	for _, v := range matrix {
		if find(v, target) {
			return true
		}
	}
	return false
}

func find(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		cmid := low + (high-low)/2
		switch {
		case nums[cmid] == target:
			return true
		case nums[cmid] < target:
			low = cmid + 1
		default:
			high = cmid - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	row, col := 0, len(matrix)-1
	for row < len(matrix[0]) && col >= 0 {
		switch {
		case matrix[col][row] == target:
			return true
		case matrix[col][row] > target:
			col--
		default:
			row++
		}
	}
	return false
}

package main

func trap1(height []int) int {
	length := len(height)
	prefix := make([]int, length)
	suffix := make([]int, length)
	// 统计前缀最大值和后缀最大值
	for i := range height {
		if i == 0 {
			prefix[i] = height[i]
			suffix[length-1-i] = height[length-1-i]
		} else {
			prefix[i] = max(height[i], prefix[i-1])
			suffix[length-1-i] = max(height[length-1-i], suffix[length-i])
		}
	}
	res := 0
	// 计算接雨水量
	for i := range height {
		res += min(prefix[i], suffix[i]) - height[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 空间优化版
func trap_twopointer(height []int) int {
	left, right := 0, len(height)-1
	if left >= right {
		return 0
	}
	leftMax, rightMax := height[left], height[right]
	res := 0
	for left < right {
		// 短板在 left ，向右移动 left
		for left < right && leftMax <= rightMax {
			left++
			if leftMax < height[left] {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
		}
		// 短板在 right ，向左移动 right
		for left < right && rightMax <= leftMax {
			right--
			if rightMax < height[right] {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
		}
	}
	return res
}

// 单调栈
func trap(height []int) int {
	stack := make([]int, 0, len(height))
	res := 0
	for i, v := range height {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if height[top] > v { // 不是 >=
				break
			}
			stack = stack[:len(stack)-1] // 出栈
			if len(stack) == 0 {
				break
			}
			top2 := stack[len(stack)-1]
			// width * height
			res += (i - top2 - 1) * (min(v, height[top2]) - height[top])
		}
		stack = append(stack, i)
	}
	return res
}

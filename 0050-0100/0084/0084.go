package main

func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	} else if l == 1 {
		return heights[0]
	}
	// 找到极小点
	minPos := findMin(heights)
	resL, resR := 0, 0
	if minPos != -1 {
		// 找到极小点后，分区寻找结果
		resL = largestRectangleArea(heights[:minPos+1])
		resR = largestRectangleArea(heights[minPos:])
	}
	res := 0
	for i, j := 0, l-1; i <= j; {
		lower := 0            // 较矮的柱子
		length := (j - i + 1) // 底宽
		// 确定较矮的柱子
		if heights[j] < heights[i] {
			lower = heights[j]
			j--
		} else {
			lower = heights[i]
			i++
		}
		// 较矮的柱子比极小值大，则停止
		if minPos != -1 && lower > heights[minPos] {
			break
		}
		// 计算当前结果
		tmp := lower * length
		if tmp > res {
			res = tmp
		}
	}
	if resL > res {
		res = resL
	}
	if resR > res {
		res = resR
	}
	return res
}

func findMin(height []int) int {
	res := -1
	up := true // 正在上升
	for j := 0; j < len(height)-1; j++ {
		if height[j] < height[j+1] {
			// 右下向上的转折点，就是极小值
			if up == false &&
				(res == -1 || height[res] > height[j]) {
				res = j
			}
			up = true
		} else if height[j] > height[j+1] {
			up = false
		}
	}
	return res
}

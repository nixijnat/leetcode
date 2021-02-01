package main

func largestRectangleArea_1(heights []int) int {
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

/////////////////////////
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	// 单调栈
	stk := make([]int, 0, 8)
	// 从左到右寻找所有左边界
	left := make([]int, n)
	for i, v := range heights {
		// 将所有大于等于当前高度的值出栈
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= v {
			stk = stk[:len(stk)-1]
		}
		if len(stk) != 0 {
			left[i] = stk[len(stk)-1]
		} else {
			left[i] = -1 // 下界
		}
		stk = append(stk, i)
	}
	stk = stk[:0]
	// 从右到左寻找所有右边界
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		// 将所有大于等于当前高度的值出栈
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) != 0 {
			right[i] = stk[len(stk)-1]
		} else {
			right[i] = n // 上界
		}
		// 入栈
		stk = append(stk, i)
	}
	res := 0
	for i, v := range heights {
		tmp := v * (right[i] - left[i] - 1)
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stk := make([]int, 0, 8)
	left := make([]int, n)
	right := make([]int, n)
	for i := range right { // 预先设置右边界
		right[i] = n
	}
	for i, v := range heights {
		stkn := len(stk)
		for ; stkn > 0 && heights[stk[stkn-1]] >= v; stkn = len(stk) {
			// 设置右边界
			right[stk[stkn-1]] = i
			stk = stk[:stkn-1]
		}
		if stkn == 0 {
			left[i] = -1
		} else {
			left[i] = stk[stkn-1]
		}
		stk = append(stk, i)
	}
	res := 0
	for i, v := range heights {
		tmp := v * (right[i] - left[i] - 1)
		if tmp > res {
			res = tmp
		}
	}
	return res
}

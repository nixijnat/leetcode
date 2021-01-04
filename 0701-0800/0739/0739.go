package main

func dailyTemperatures(T []int) []int {
	length := len(T)
	res := make([]int, length)
	for i := 0; i < length; i++ {
		tmp := 0
		for j := i + 1; j < length; j++ {
			if T[i] < T[j] {
				tmp = j - i
				break
			}
		}
		res[i] = tmp
	}
	return res
}

func dailyTemperatures2(T []int) []int {
	length := len(T)
	res := make([]int, length)
	stack := make([]int, length) // 存储索引
	for i := 0; i < length; i++ {
		// 从栈中不断取元素比较元素，如果
		for l := len(stack); l > 0 && T[stack[l-1]] < T[i]; l-- {
			res[stack[l-1]] = i - stack[l-1]
			stack = stack[:l-1]
		}
		// 把当前元素放入栈中
		stack = append(stack, i)
	}
	return res
}

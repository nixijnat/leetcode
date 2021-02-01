package main

func leastInterval(tasks []byte, n int) int {
	// 记录数量
	m := [26]int{}
	max1 := 0 // 任务最大数量
	for _, v := range tasks {
		m[v-'A']++
		max1 = max(max1, m[v-'A'])
	}
	maxNum := 0 // 任务最多的任务类型数量
	for _, v := range m {
		if v == max1 {
			maxNum++
		}
	}
	return max(len(tasks), (max1-1)*(n+1)+maxNum)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

func subarraySum(nums []int, k int) int {
	prefixSumMap := make(map[int]int, len(nums))
	prefixSumMap[0] = 1 // 必须默认加起始位
	res := 0
	prefixSum := 0
	for _, v := range nums {
		// 当前值和前缀和
		prefixSum += v
		res += prefixSumMap[prefixSum-k]
		// 存储当前值和前缀和：如果求范围，则需要注意
		prefixSumMap[prefixSum]++
	}
	return res
}

// 存储前缀和
func subarraySum2(nums []int, k int) int {
	prefixSumMap := make(map[int]int, len(nums))
	res := 0
	prefixSum := 0
	for _, v := range nums {
		// 记录下这个前缀，如果需要具体的范围，可以记录索引
		prefixSumMap[prefixSum]++
		prefixSum += v // 作为下一个数的前缀和
		res += prefixSumMap[prefixSum-k]
	}
	return res
}

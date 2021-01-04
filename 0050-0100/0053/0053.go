package main

func maxSubArray1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	res, tmp := nums[0], 0
	for _, v := range nums {
		tmp += v        // 和
		if tmp >= res { // 更新 结果
			res = tmp
		}
		if tmp < 0 { // 前面和已经为负数 则丢弃
			tmp = 0
		}
	}
	return res
}

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	res, cur := nums[0], 0
	for _, v := range nums {
		cur += v // f(i-1) + v
		if cur < v {
			cur = v // max (v,  f(i-1) + v )
		}
		if res < cur { // 前面和已经为负数 则丢弃
			res = cur
		}
	}
	return res
}

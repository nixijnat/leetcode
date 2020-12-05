package main

func tryMakeByDivisor(target, divisor) []int {
	tmp := target % divisor // 是否能被除尽
	if tmp != 0 {
		return nil
	}
	tmp = target / divisor // 数量
	ret := make([]int, 0, tmp)
	for i := 0; i < l; i++ {
		ret = append(ret, divisor)
	}
	return ret
}

func combinationSum(nums []int, target int) [][]int {
	l := len(nums)
	if l == 0 {
		return nil
	} else if l == 1 {
		// 只有一个元素时判断是否是约数
		return tryMakeByDivisor(target, nums[0])
	}
	ret := make([][]int, 0, 4)
	for i := range nums {
		if target == nums[i] {
			ret = append(ret, []int{target})
			continue
		} else if target < nums[i] {
			continue
		}
		// 递归找余值：以 nums[i] 为基准，在 [i:] 的范围内寻找
		tmp := combinationSum(nums[i:], target-nums[i])
		if len(tmp) == 0 {
			continue
		}
		for _, v := range tmp {
			ret = append(ret, append(v, nums[i])) // 附加上基准值
		}
	}
	return ret
}

package main

func moveZeroes(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	p0 := 0 // 指向 0
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			// i 指向0时，p0指向非0，则把 p0 移到 i
			if nums[p0] != 0 {
				p0 = i
			}
		} else {
			// i 指向非0时，与 p0 交换数
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

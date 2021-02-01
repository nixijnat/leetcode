package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	switch {
	case l1 == 0 && l2 == 0:
		return 0
	case l1 == 0 && l2 == 1:
		return float64(nums2[0])
	case l1 == 1 && l2 == 0:
		return float64(nums1[0])
	}
	// 比较两个数大小，处理了边界问题
	less := func(i1, i2 int) bool {
		switch {
		case i1 >= l1:
			return false
		case i2 >= l2:
			return true
		default:
			return nums1[i1] < nums2[i2]
		}
	}
	remain := (l1 + l2) >> 1 // N/2 除2 ，找的第 remain 个元素
	s1, s2 := 0, 0           // 记录每次新数组的起始位置
	for remain > 1 {         // 从第一个元素找第N个元素，应该走N-1步
		half := remain >> 1
		if less(s1+half-1, s2+half-1) {
			s1 += half
		} else {
			s2 += half
		}
		remain -= half
	}
	// 得出较小的数，处理了边界问题
	min := func(i1, i2 int) int {
		if less(i1, i2) {
			return nums1[i1]
		}
		return nums2[i2]
	}
	k, k1 := 0, 0
	if less(s1, s2) {
		k, k1 = nums1[s1], min(s1+1, s2)
	} else {
		k, k1 = nums2[s2], min(s1, s2+1)
	}
	if (l1+l2)&0x1 != 0 { // N%2取余的快速法
		return float64(k1)
	}
	return float64(k+k1) / 2
}

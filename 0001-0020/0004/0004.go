package main

func findMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 == 0 && l2 == 0 {
		return 0
	} else if l1 == 0 && l2 == 1 {
		return float64(nums2[0])
	} else if l1 == 1 && l2 == 0 {
		return float64(nums1[0])
	}
	// 比较两个数大小，处理了边界问题
	less := func(i1, i2 int) bool {
		if i1 >= l1 {
			return false
		} else if i2 >= l2 {
			return true
		} else {
			return nums1[i1] < nums2[i2]
		}
	}
	// 求二分后的索引长度
	halfIndex := func(len int) int {
		if len <= 1 {
			return 0
		}
		return len>>1 - 1 // 长度/2-1
	}
	targetNum := (l1 + l2) >> 1 // N/2 除2 ，找的第 target 个元素
	s1, s2 := 0, 0              // 记录每次新数组的起始位置
	i, j := halfIndex(targetNum), halfIndex(targetNum)
	for targetNum > 1 { // 从第一个元素找第N个元素，应该走N-1步
		if less(i, j) {
			s1 = i + 1
		} else {
			s2 = j + 1
		}
		targetNum -= targetNum / 2
		i, j = s1+halfIndex(targetNum), s2+halfIndex(targetNum)
	}
	// 得出较小的数，处理了边界问题
	min := func(i1, i2 int) int {
		if i1 >= l1 {
			return nums2[i2]
		} else if i2 >= l2 {
			return nums1[i1]
		} else {
			if nums1[i1] < nums2[i2] {
				return nums1[i1]
			} else {
				return nums2[i2]
			}
		}
	}
	target1, target2 := 0, 0
	if less(i, j) {
		target1 = nums1[i]
		target2 = min(i+1, j)
	} else {
		target1 = nums2[j]
		target2 = min(i, j+1)
	}
	if (l1+l2)&0x01 != 0 { // N%2取余的快速法
		return float64(target2)
	} else {
		return float64(target1+target2) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}

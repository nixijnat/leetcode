package main

import (
	"sort"
)

func main() {

}

func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	} else if l == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{nums}
		} else {
			return nil
		}
	}
	m := make(map[int]int, l)
	for _, v1 := range nums {
		m[v1]++
	}
	ns := nums[:0]
	for v := range m {
		ns = append(ns, v)
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i] < ns[j]
	})
	ret := make([][]int, 0, l)
	for i, j := 0, len(ns)-1; i <= j; {
		if ns[i] > 0 || ns[j] < 0 {
			break
		}
		if i == j {
			if m[0] >= 3 {
				ret = append(ret, []int{0, 0, 0})
			}
			break
		}
		v3 := 0 - ns[i] - ns[j]
		count, ok := m[v3]
		//fmt.Println(ns[i], ns[j], v3, count)
		if ok && !((v3 == ns[i] || v3 == ns[j]) && count <= 1) {
			ret = append(ret, []int{ns[i], ns[j], v3})
		}
		if v3 < 0 { // ns[i] + ns[j] >= 0
			delete(m, ns[j])
			j--
		} else {
			delete(m, ns[i])
			i++
		}
	}
	return ret
}

func threeSum2(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	} else if l == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{nums}
		} else {
			return nil
		}
	}
	sort.Ints(nums) // 有专门处理 整型排序的
	ret := make([][]int, 0, l/3)
	for i := 0; i < l-2 && nums[i] <= 0; i++ {
		for ; i > 0 && i < l && nums[i] == nums[i-1]; i++ {
		}
		for j, k := i+1, l-1; j < k; {
			if nums[k] < 0 { // 剔除掉大于0的可能
				break
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
			}
			if sum <= 0 {
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
			}
			if sum >= 0 {
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}
	return ret
}

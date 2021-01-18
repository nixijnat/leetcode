package main

import "sort"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	sort.Slice(coins, func(i, j int) bool { return coins[i] > coins[j] })
	var find func(coins []int, amount int, count, res int) int
	find = func(coins []int, amount int, count, res int) int {
		if amount == 0 {
			return min(count, res)
		}
		ok := false
		for i, v := range coins {
			if amount >= v && count < res {
				l := find(coins[i:], amount-v, count+1, res)
				if l != -1 {
					ok = true
					res = min(l, res)
				}
			}
		}
		if !ok {
			return -1
		}
		return res
	}
	return find(coins, amount, 0, amount)
}

package main

import (
	"testing"
)

func TestTowSum(t *testing.T) {
	type m struct {
		Input1 []int
		Input2 int
		Expect []int
	}

	t.Run("sort", func(t *testing.T) {
		var ts = []m{
			{[]int{2, 11, 7, 15}, 9, []int{0, 2}},
			{[]int{2, 11, 7, 15}, 13, []int{0, 1}},
			{[]int{2, 11, 7, 15}, 18, []int{1, 2}},
			{[]int{2, 11, 7, 15}, 26, []int{1, 3}},
			{[]int{3, 2, 4}, 6, []int{1, 2}},
			{[]int{230, 863, 916, 585, 981, 404, 316, 785, 88, 12, 70, 435, 384, 778, 887, 755, 740, 337, 86, 92, 325, 422, 815, 650, 920, 125, 277, 336, 221, 847, 168, 23, 677, 61, 400, 136, 874, 363, 394, 199, 863, 997, 794, 587, 124, 321, 212, 957, 764, 173, 314, 422, 927, 783, 930, 282, 306, 506, 44, 926, 691, 568, 68, 730, 933, 737, 531, 180, 414, 751, 28, 546, 60, 371, 493, 370, 527, 387, 43, 541, 13, 457, 328, 227, 652, 365, 430, 803, 59, 858, 538, 427, 583, 368, 375, 173, 809, 896, 370, 789}, 542, []int{28, 45}},
		}
		for _, i := range ts {
			tmp := twoSumSort(i.Input1, i.Input2)
			if tmp == nil || tmp[0] != i.Expect[0] || tmp[1] != i.Expect[1] {
				t.Error(i, tmp)
			}
		}
	})
	t.Run("map", func(t *testing.T) {
		var ts = []m{
			{[]int{2, 11, 7, 15}, 9, []int{0, 2}},
			{[]int{2, 11, 7, 15}, 13, []int{0, 1}},
			{[]int{2, 11, 7, 15}, 18, []int{1, 2}},
			{[]int{2, 11, 7, 15}, 26, []int{1, 3}},
			{[]int{3, 2, 4}, 6, []int{1, 2}},
			{[]int{230, 863, 916, 585, 981, 404, 316, 785, 88, 12, 70, 435, 384, 778, 887, 755, 740, 337, 86, 92, 325, 422, 815, 650, 920, 125, 277, 336, 221, 847, 168, 23, 677, 61, 400, 136, 874, 363, 394, 199, 863, 997, 794, 587, 124, 321, 212, 957, 764, 173, 314, 422, 927, 783, 930, 282, 306, 506, 44, 926, 691, 568, 68, 730, 933, 737, 531, 180, 414, 751, 28, 546, 60, 371, 493, 370, 527, 387, 43, 541, 13, 457, 328, 227, 652, 365, 430, 803, 59, 858, 538, 427, 583, 368, 375, 173, 809, 896, 370, 789}, 542, []int{28, 45}},
		}
		for _, i := range ts {
			tmp := twoSumMap(i.Input1, i.Input2)
			if tmp == nil || tmp[0] != i.Expect[0] || tmp[1] != i.Expect[1] {
				t.Error(i, tmp)
			}
		}
	})
}

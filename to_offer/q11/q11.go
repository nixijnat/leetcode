package q11

// 11. 旋转数组的最小数字

// 5 种情况 / 和 - 表示坡度 v表示坑
// 1. i////mid//V//j  取右
// 2. i//V//mid////j  取左
// 3. i--V--mid----j  无法判断 i++
// 4. i----mid--V--j  无法判断 i++
// 5. i////mid////j  取左

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	var i int
	for j := len(numbers) - 1; i < j; {
		mid := i + (j-i)>>1
		switch {
		case numbers[mid] > numbers[j]:
			i = mid + 1
		case numbers[mid] < numbers[i] || numbers[i] < numbers[j]:
			j = mid
		default:
			i++
		}
	}
	return numbers[i]
}

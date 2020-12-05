package main

func main() {

}

func removeDuplicates1(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	tmp := nums[:1]
	for i := 1; i < length; i++ {
		if nums[i-1] != nums[i] {
			tmp = append(tmp, nums[i])
		}
	}
	return len(tmp)
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	slow, fast := 1, 1 // slow 是下一个放元素的位置
	for ; fast < length; fast++ {
		if nums[fast-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

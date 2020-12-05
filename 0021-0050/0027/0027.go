package main

func main() {

}

func removeElement1(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	tmp := nums[:0]
	for i := 0; i < length; i++ {
		if val != nums[i] {
			tmp = append(tmp, nums[i])
		}
	}
	return len(tmp)
}

func removeElement2(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	slow, fast := 0, 0 // slow 是下一个放元素的位置
	for ; fast < length; fast++ {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	for i := 0; i < length; {
		if val == nums[i] {
			length--
			nums[i] = nums[length]
		} else {
			i++
		}
	}
	return length
}

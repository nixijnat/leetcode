package main

func intToRoman1(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	reps := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	buf := ""
	for i := 0; i < 13; i++ {
		for num >= values[i] {
			num -= values[i]
			buf += reps[i]
		}
	}
	return buf
}
func intToRoman(num int) string {
	if num <= 0 {
		return ""
	}
	nums := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	buf := ""
	for i := 0; i < 13; {
		if num >= nums[i].value {
			num -= nums[i].value
			buf += nums[i].symbol
		} else {
			i++
		}
	}
	return buf
}

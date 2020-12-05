package main

func isPalindrome(x int) bool {
	switch {
	case x < 0:
		return false
	case x < 10:
		return true
	case x%10 == 0:
		return false
	}
	newNum := 0
	for i := x; i != 0; i /= 10 {
		newNum = newNum*10 + i%10
	}
	return newNum == x
}

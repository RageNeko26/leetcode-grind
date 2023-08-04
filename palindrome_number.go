package main

func isPalindrome(x int) bool {
	var reversedNum int
	tmp := x

	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}

	if x == reversedNum {
		return true
	}

	return false

}

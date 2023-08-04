package main

import "fmt"

func reverse(x int) int {
	var revNum int
	tmp := x

	for tmp > 0 {
		revNum = revNum*10 + tmp%10
		tmp = tmp / 10
	}

	return revNum
}

func main() {
	result := reverse(123)
	fmt.Println(result)
}

package main

import "fmt"

func main() {
	num := 5
	f := factorial(num)
	fmt.Println("factporial:", f)
}

func factorial(num int) int {
	println(num)
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// 5 * factocial(4) * factocial(3) factocial(2) * factocial(1) *1

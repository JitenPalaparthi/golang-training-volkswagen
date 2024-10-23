package main

import "fmt"

func main() {

	func() { // body of the function

	}() // () is an executor

	func() {
		fmt.Println("Hello World!")
	}()

	func(msg string) {
		fmt.Println(msg)
	}("Hello VolksWagen folks!")

	func(nums ...int) {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		println("Sum:", sum)

	}(1, 3, 5, 3, 5, 4, 7, 57, 34, 3, 2, 67, 67)

	sum1 := func(nums ...int) int {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		return sum
	}(1, 3, 5, 3, 5, 4, 7, 57, 34, 3, 2, 67, 67)
	println("Sum:", sum1)

	sumf := func(nums ...int) int {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		return sum
	}
	sum2 := sumf(1, 3, 5, 3, 5, 4, 7, 57, 34, 3, 2, 67, 67)
	println("Sum:", sum2)
}

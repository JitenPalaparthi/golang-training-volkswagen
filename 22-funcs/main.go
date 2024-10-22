package main

import "fmt"

func main() {

	s1 := sumOfVariadic()
	fmt.Println("s1:", s1)

	s2 := sumOfVariadic(10, 20)
	fmt.Println("s2:", s2)

	s3 := sumOfVariadic(10, 20, 23, 43, 45, 45, 6, 5, 56, 67)
	fmt.Println("s3:", s3)

	s4 := sumOfMulby(2, 2, 4, 6, 8)
	fmt.Println("s4:", s4)
}

func sumOf(a, b int) int {
	return a + b
}

func sumOfSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumOfVariadic(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func sumOfMulby(nums ...int, num int) int { // this is not okay
func sumOfMulby(num int, nums ...int) int { // this is okay
	sum := 0
	for _, v := range nums {
		sum += (v * num)
	}
	return sum
}

// variadic paramter can pass any number of arguments of same type
// the variadic parameter must be the last parameter in a function
// variadic parameter can only be used in functions and methods

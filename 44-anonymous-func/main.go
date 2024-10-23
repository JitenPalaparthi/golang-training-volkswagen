package main

import "fmt"

func main() {
	var fn1 func(int, int) int
	//fmt.Println(Fn1)
	// add
	a, b := 10, 20
	r1 := Calc(a, b, func(a1, b1 int) int {
		return a1 + b1
	})
	fmt.Println("Add:", r1)

	// sub
	fn1 = func(i1, i2 int) int {
		return i1 - i2
	}
	r2 := Calc(a, b, fn1)
	fmt.Println("Sub:", r2)

	//mul

	r3 := Calc(a, b, Mul)
	fmt.Println("Mul:", r3)

}

func Calc(a, b int, fn1 func(int, int) int) int {
	return fn1(a, b) // function is executed and returned
}

func Mul(a, b int) int {
	return a * b
}

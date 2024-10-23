package main

import "fmt"

var Result int

type MyFunc func(int, int) int

func (m *MyFunc) Display() func() {
	fmt.Println("The result is:", Result)
	return func() {
		fmt.Println("Calling a return function")
	}
}

func main() {
	var fn1 MyFunc = func(i1, i2 int) int {
		Result = i1 + i2
		return Result
	}

	fmt.Println(fn1(10, 20))
	fn2 := fn1.Display()
	fn2()
}

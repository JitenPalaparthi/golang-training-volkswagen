package main

import (
	"runtime"
)

func main() {
	println("Main has started")
	go func(num int) {
		for i := 1; i <= num; i++ {
			if i%2 != 0 {
				println("Odd", i)
			}
		}
	}(300)
	go printEven(200)

	println("Main is going to be exited")
	runtime.Goexit() // it does not return exit code 0 on main
}

func printEven(num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			println("Even", i)
		}
	}
}

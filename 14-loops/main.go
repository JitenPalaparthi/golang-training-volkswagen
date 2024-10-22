package main

import "fmt"

func main() {

	// loop with multiple initializers
	for i, j := 0, 10; i < j && true && true && (true || false); i, j = i+1, j-1 {
		fmt.Printf(" i:%d j:%d\n", i, j)
	}

	println()
	println("break only inner")
	for i := 0; i <= 10; i++ {
		for j := 10; j >= 1; j-- {
			if i == j {
				break
			}
			fmt.Printf(" i:%d j:%d\n", i, j)
		}
	}
	println()

	println("break outer")
outer:
	for i := 0; i <= 10; i++ {
		for j := 10; j >= 1; j-- {
			if i == j {
				break outer
			}
			fmt.Printf(" i:%d j:%d\n", i, j)
		}
	}
	println()

	println("break manually with flag")
	ok := false
	for i := 0; i <= 10; i++ {
		if ok {
			break
		}
		for j := 10; j >= 1; j-- {
			if i == j {
				ok = true
				break
			}
			fmt.Printf(" i:%d j:%d\n", i, j)
		}
	}

}

package main

func main() {
	defer func() {
		defer println("Second defer")
	}()
	defer println("\nSecond defer")

	println("hello i am main starting")

	str := "Hello World!"

	for _, v := range str {
		defer print(string(v)) //
		// !dlrow olleH
	}

	/*
		!
		d
		l
		r
		o
		W

		o
		l
		l
		e
		H

	*/

	println()
}

// defer defers the execution to the end

// defer calls are there in the stack

// the last defer in the call stack is executed at first

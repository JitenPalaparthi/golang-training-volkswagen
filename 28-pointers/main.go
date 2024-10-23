package main

import (
	"errors"
)

func main() {

	var num1 int = 100

	var ptr1 *int = &num1

	// var ptr2 **int = &ptr1

	// var ptr3 ***int = &ptr2

	*ptr1 = 200 // derefecencing a pointer

	println(num1)

	changVal(num1)

	println(num1) // what is the value of num1
	changValPtr(ptr1)
	println(num1) // what is the value of num1
	changValPtr(&num1)
	println(num1) // what is the value of num1

	var ptr2 *int            // default value nil
	err := changValPtr(ptr2) // trying to dereference a pointer which is nil

	if err != nil {
		println(err.Error())
	} else {
		println(*ptr2)
	}

}

// pointer holds the address
// dereference a pointer
// nil pointer
// can check nil for a pointer

func changVal(num int) { // copy
	num++
}

func changValPtr(num *int) error { // reference
	if num != nil {
		*num++
		return nil
	}
	return errors.New("nil pointer")
}

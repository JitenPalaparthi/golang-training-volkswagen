package main

import "unsafe"

func main() {
	//println(ptr2, *ptr2)
	arr1 := [3]int{10, 20, 30}
	//arr2 := [3]bool{true, false, true}
	var address1 uintptr = uintptr(unsafe.Pointer(&arr1[0]))
	address1 += unsafe.Sizeof(arr1[0])
	val := (*int)(unsafe.Pointer(address1))
	println(*val)

	ptr1 := new(int) // new allocates the memory based on the type and returns the pointer
	println(*ptr1)

	*ptr1 = 100
	println(*ptr1)

	//var ptr2 *int // it is a nil pointer until an address is assigned

	ptr3 := new(string) // this is not a nil pointer as address is assigned and a value is inferred
	println(*ptr3)
}

package main

import "fmt"

func main() {

	// short num1 = 1312312
	// long num2 = num1 // implicit casting

	var num1 uint8 = 213

	//var num2 uint64 = num1 // no automatic casting done here

	var num2 uint64 = uint64(num1) // type casting
	println(num1, num2)

	float1 := 231231.12312
	num2 = uint64(float1)

	println(float1, num2)

	var num3 uint64 = 12321312112 // 8 bytes
	// 1011011110011010000100110101110000
	fmt.Printf("num3:%d binary of num3:%b\n", num3, num3)
	var num4 uint8 = uint8(num3)   // cast to 1 byte
	var num5 uint16 = uint16(num3) // cast to 1 byte
	//01110000 -> 112
	// 0100110101110000 -> 19824
	println(num4, num5)
	// ok1 := true
	// num1 =uint8(ok1)
}

// casting
//

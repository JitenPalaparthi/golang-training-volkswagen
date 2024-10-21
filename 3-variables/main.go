package main

import "fmt"

var ( // global variable
	maxnum int = 9999999999
)

const DAYS int = 7 // constant

const DAYHOURS = 24

func main() {

	// all these are local variables

	var num1 int = 1231231

	var num2 uint8 = 255

	var (
		float1 float32 = 123.123

		float2 float64 = 12312.12312

		ok1 bool = true

		str1 string = "VolksWagen"
	)

	var num3, num4 uint16 = 213, 123

	println(num1, num2, num3, num4, float1, float2, ok1, str1)

	const PI float32 = 3.14

	// type inference
	var num5 = 12321

	var age1 = 39 //int 8 bytes on 64 bit machine

	var age2 uint8 = 39

	var value1 = 12.34 // float64

	var value2 float32 = 12.34

	var ok2 = true

	var str2 = "Jiten Palaparthi"

	var (
		num6, num7 = 100, 200
		ok3, ok4   = true, false
	)

	// value inference for a type

	var ok5 bool    // value inference for the type bool is false
	var num8 uint8  // value inference for the number is 0
	var str3 string // value inference for string is "" //empty string

	println(num5, age1, age2, value1, value2, ok2, str2, num6, num7, ok3, ok4, ok5, num8, str3)
	// properformat of float

	fmt.Printf("value1:%.2f value2:%.4f\n", value1, value2)
}

/*

numbers:
	uint,int, uint8,uint16,uint32,uint64, int8,int16,int32,int64,float32,float64
	uintptr, rune, byte complex64,complex128

boolean:
	bool

strings:
	string

interface:
	interface{} or any

*/

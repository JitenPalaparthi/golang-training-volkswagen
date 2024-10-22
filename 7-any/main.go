package main

import (
	"fmt"
	"reflect"
)

func main() {
	var iface1 any
	//var iface2 interface{}

	println(iface1)
	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

	iface1 = true
	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

	iface1 = 12312.12312
	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

	iface1 = 12312
	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

	iface1 = "Hello world"
	fmt.Println("Value:", iface1, "Type:", reflect.TypeOf(iface1))

	// type casting in any?

	var (
		num1   int
		float1 float32
		float2 float64
		any1   any = 13123131 // what is the type of any1 here ?
	)

	num1 = any1.(int) // any.(type) .. canot do type cast but only type assertion bcz of any1 is any

	fmt.Println(num1, "Type:", reflect.TypeOf(any1))

	any1 = 1232.12312 // float64

	float1 = float32(any1.(float64))
	fmt.Println(float1, "Type:", reflect.TypeOf(any1))

	any1 = 312321
	float2, ok := any1.(float64)

	if ok {
		fmt.Println(float2, "Type:", reflect.TypeOf(any1))
	} else {
		num4, ok1 := any1.(int)
		if ok1 {
			fmt.Println(num4, "Type:", reflect.TypeOf(num4))
		} else {
			fmt.Println("cound not be asserted either to float or to int")
		}
	}
}

// interface{} --> empty interface
// any can hold any type and so its data
// the default value and type of any are nil and nil respectively
// the type of any variable is based on the value that any has at that moment
// cannot type cast any rather do type assertion

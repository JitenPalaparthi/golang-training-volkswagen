package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {

	//var num int = 100

	var arr1 [5]int
	var arr2 [5]int
	var arr3 [4]int

	fmt.Println("arr1:", arr1, "type:", reflect.TypeOf(arr1))

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Intn(100)
	}
	fmt.Println(arr1)

	// sum of elements in an array

	sum := 0

	for _, v := range arr1 {
		sum += v
	}
	println("Sum of arr1:", sum)
	arr2 = arr1 // both arrays are same type
	//arr3 = arr2 // cant work bcz both of them are two different types

	// copy elements of an array when both of them are two different types w.r.t length
	l := min(len(arr2), len(arr3))

	for i := 0; i < l; i++ {
		arr3[i] = arr2[i]
	}
	fmt.Println("arr1:", arr1, "\narr2:", arr2, "\narr3:", arr3)
}

// array is a collection of elements those belong to the same data type
// array size/ length should be aware at the compile time.
// array cannot be grown or shrunk at runtime
// there is type and value inferrence to the array

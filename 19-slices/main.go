package main

import (
	"fmt"

	"math/rand"
)

func main() {

	var slice1 []int // a slice is declared but not instantiated

	if slice1 == nil {
		println("slice is not instantiated")
		slice1 = make([]int, 5, 10) // instantiatig a slice with length , zeros are inferred to all 5 elements of the slice
	}
	fmt.Println(slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}

	fmt.Println(slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

	for i := 1; i <= 5; i++ {
		slice1 = append(slice1, rand.Intn(100))
	}
	fmt.Println(slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

	slice1 = append(slice1, rand.Intn(100))
	fmt.Println(slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

	for i := 1; i <= 10; i++ {
		slice1 = append(slice1, rand.Intn(100))
	}
	fmt.Println(slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

}

// make is a built in function
// used to crete slice, map and channel
// slice gives reference when instantiated
// slice can grow or shrink at runtime

/* slice struct
ptr *int // 8
len int  // 8
cap int  // 8
*/

/*
slice1
ptr: 0x14000136050
len: 5
cap: 10
*/

/*
slice1
ptr: 0x14000136050
len: 10
cap: 10
*/

/*
slice1
ptr: 0x1400013e000
len: 11
cap: 20
*/

/*
slice1
ptr: 0x1400013a000
len: 21
cap: 40
*/

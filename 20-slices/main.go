package main

import "fmt"

func main() {

	slice1 := []int{10, 20} // init with values
	fmt.Println("slice1", slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

	slice2 := make([]int, 2, 2)
	slice2[0] = 123
	slice2[1] = 423
	fmt.Println("slice2", slice2, "len:", len(slice2), "cap:", cap(slice2), &slice2[0])

	var slice3 []int

	// append function ,if the slice is nill then append function instntiate the slice and append values to it
	slice3 = append(slice3, 30, 40) // variadic parameter
	fmt.Println("slice3", slice3, "len:", len(slice3), "cap:", cap(slice3), &slice3[0])

	slice4 := slice3 // shallow copy
	fmt.Println("slice4", slice4, "len:", len(slice4), "cap:", cap(slice4), &slice4[0])

	slice4 = append(slice4, 400)
	slice3 = slice4
	slice4[1] = 1232331
	fmt.Println("slice3", slice3, "len:", len(slice3), "cap:", cap(slice3), &slice3[0])
	fmt.Println("slice4", slice4, "len:", len(slice4), "cap:", cap(slice4), &slice4[0])

}

/*
slice3
ptr: 0x1400013a000
len: 2
cap: 3
*/

/*
slice4
ptr: 0x1400013a000
len: 2
cap: 3
*/

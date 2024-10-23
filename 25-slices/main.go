package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 5)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(100)
	}

	slice2 := slice1 // a shallow copy

	slice3 := make([]int, 5)

	copy(slice3, slice1) // deep copy or clone
	// if the length of the destication is smaller than the source then only des len elements are copied
	// if dest length is bigger , then all elements of source are copied the rest of the elements aregiven zeros

	fmt.Println(slice1, slice2, slice3)

	clear(slice2)
	// it clears the values of a slice and defaults all the elements of the slice
	// since slice2 is a shallow copy of slice1 both slice2 and slice1 are cleared
	fmt.Println(slice1, slice2, slice3)

}

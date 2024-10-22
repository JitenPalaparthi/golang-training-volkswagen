package main

import (
	"errors"
	"fmt"
)

func main() {

	var slice1 []int
	slice1 = append(slice1, 10, 20, 30, 40, 50)
	fmt.Println("slice1", slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

	arr1 := [5]int{10, 20, 30, 40, 50}

	slice2 := slice1      // whole slice is shallwo copied to the another slice
	slice3 := slice1[:]   // all elements
	slice4 := slice1[3:]  // from 3rd to the end
	slice5 := slice1[:3]  // from 0th to the 3rd but not 3rd
	slice6 := slice1[2:4] // starts from 2,3 but 4th is not included

	slice7 := arr1[:] // all elements of arr1

	slice8 := arr1[2:4] // starts from 2,3 but 4th is not included

	fmt.Println("slice2", slice2, "len:", len(slice2), "cap:", cap(slice2), &slice2[0])
	fmt.Println("slice3", slice3, "len:", len(slice3), "cap:", cap(slice3), &slice3[0])
	fmt.Println("slice4", slice4, "len:", len(slice4), "cap:", cap(slice4), &slice4[0])
	fmt.Println("slice5", slice5, "len:", len(slice5), "cap:", cap(slice5), &slice5[0])
	fmt.Println("slice6", slice6, "len:", len(slice6), "cap:", cap(slice6), &slice6[0])

	fmt.Println("slice7", slice7, "len:", len(slice7), "cap:", cap(slice7), &slice7[0])
	fmt.Println("slice8", slice8, "len:", len(slice8), "cap:", cap(slice8), &slice8[0])
	fmt.Println("arr1", arr1, "len:", len(arr1), "cap:", cap(arr1), &arr1[0])
	slice8[0] = 3000 // though slice is changed , the impact will be on array since it is a shallow copy
	fmt.Println("arr1", arr1, "len:", len(arr1), "cap:", cap(arr1), &arr1[0])

	sum1 := sumOfVariadic(slice1...)
	sum2 := sumOfVariadic(arr1[:]...)
	fmt.Println("Sum1:", sum1)
	fmt.Println("Sum2:", sum2)

	// delete a value from a slice
	//slice1 = append(slice1[:2], slice1[3:]...)

	slice1 = removeElemSlice(0, slice1)

	// indexes --> 0,1,3,4
	fmt.Println("slice1", slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])
	slice1 = removeElemSlice(0, slice1)

	// indexes --> 0,1,3,4
	fmt.Println("slice1", slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

	slice1 = removeElemSlice(0, slice1)

	// indexes --> 0,1,3,4
	fmt.Println("slice1", slice1, "len:", len(slice1), "cap:", cap(slice1), &slice1[0])

	var slice9 []int
	slice9, err := removeElem(0, slice9)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("slice9", slice9, "len:", len(slice9), "cap:", cap(slice9), &slice9[0])

	}

	if slice10, err := removeElem(10, arr1[:]); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("slice10", slice10, "len:", len(slice10), "cap:", cap(slice10), &slice10[0])
	}

}

func sumOfVariadic(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func removeElemSlice(index int, slice []int) []int {

	if slice == nil {
		return slice
	}
	if index+1 >= len(slice) {
		return slice
	}

	slice = append(slice[0:index], slice[index+1:]...)

	return slice
}

func removeElem(index int, slice []int) ([]int, error) {
	if slice == nil {
		return nil, errors.New("input slice is nil")
	}
	if index+1 >= len(slice) {
		return nil, fmt.Errorf("index is beyond the range.Max index is %d", len(slice)-1)
	}
	slice = append(slice[0:index], slice[index+1:]...)
	return slice, nil
}

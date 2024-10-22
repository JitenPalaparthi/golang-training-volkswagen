package main

import "fmt"

func main() {

	// shorthand declaraion of array

	arr1 := [5]int{10, 20, 30, 40, 50}

	fmt.Println(arr1)

	arr2 := [...]int{12, 433, 345, 45, 6, 5, 667} // the length is evaluated at compile time

	fmt.Println(arr2)

	arr3 := [...]any{"Hello", 123, 4343.343, true, 'A'}

	fmt.Println(arr3)

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2d)

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			println(arr2d[i][j])
		}
	}

	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	fmt.Println(arr3d)

	for _, arr1 := range arr3d {
		for _, arr2 := range arr1 {
			for _, v := range arr2 {
				println(v)
			}
		}
	}

}

package main

import "fmt"

func main() {
	fmt.Println("hello VolksWagen folks")
	greet()
	a1 := areaOfRect(12.34, 15.34)
	fmt.Printf("Area of rect a1:%.3f\n", a1)
	p1 := perimeterOfRect(12.34, 15.34)
	fmt.Printf("Perimeter of rect p1:%.3f\n", p1)

	arr1 := [...]int{10, 20, 30, 40, 50}
	sum1 := sumOfArray(arr1)

	fmt.Println("Sum of arr1:", sum1)

	arr2 := [...]int{10, 20, 30, 40}
	// sum2 := sumOfArray(arr2)

	// fmt.Println("Sum of arr2:", sum2)
	//slice1 := arr2[:] // converting and assigning an array as a slice
	sum2 := sumOf(arr2[:])
	fmt.Println("Sum of arr2:", sum2)

	slice2 := reverse(arr2[:])
	fmt.Println("slice2:", slice2)

	a3, p3 := areaAndPerimeterOfRect(12.34, 15.34)
	fmt.Printf("Area a3:%.3f perimeter: %.3f\n", a3, p3)

	a4, _ := areaAndPerimeterOfRect(12.34, 15.34)
	fmt.Printf("Area a4:%.3f \n", a4)
	_, p4 := areaAndPerimeterOfRect(12.34, 15.34)
	fmt.Printf("Petimeter p4:%.3f \n", p4)
}

func greet() {
	println("hello world")
}

func areaOfRect(l, b float32) float32 {
	return l * b
}

func perimeterOfRect(l float32, b float32) float64 {
	return float64(2.0 * (l + b))
}

func areaAndPerimeterOfRect(l, b float32) (a float32, p float32) {
	a, p = (l * b), 2*(l+b)
	return a, p
}

func sumOfArray(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func sumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func reverse(slice []int) (retSlice []int) {
	retSlice = make([]int, len(slice))

	for i, j := len(slice)-1, 0; i >= 0; i, j = i-1, j+1 {
		retSlice[j] = slice[i]
	}

	return retSlice
}

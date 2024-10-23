package main

import "math/rand"

func main() {
	r1 := stackAllocation(100)
	println("r1", r1, "address of r1:", &r1)

	r2 := heapAllocation(100)
	println("r2", *r2, "address of r2:", r2)

	m1 := map[string]any{"1": 100, "2": 200}
	println(m1)

	slice1 := []int{1, 2, 3}
	println(slice1)

	for i := 1; i <= 1000000; i++ {
		slice1 = append(slice1, rand.Intn(12313))
	}
	println(slice1, len(slice1))

	slice2 := make([]int, 1000000)
	println(slice2)

	for i := 1; i <= 100; i++ {
		slice2 = append(slice2, rand.Intn(12313))
	}
	println(slice2, len(slice2))

	var any1 any = 1000
	any1 = "hello world"
	{
		println(any1)
		any1 = m1
		println(any1)
	}

}

// escape anlaysis

func stackAllocation(num int) int {
	x := num * num // new variable
	println("x", x, "address of x:", &x)
	return x
}

func heapAllocation(num int) *int { // generaly this is an error code in c,c++
	x := new(int)
	*x = num * num // new variable
	println("x", *x, "address of x:", x)
	return x
} // but go escape analysis handles this in a sepate way

// how to do escape analysis
// go build -gcflags="-m" main.go

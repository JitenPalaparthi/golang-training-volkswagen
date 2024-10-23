package main

import (
	"fmt"
	"shapes-demo/shape"
	"shapes-demo/shape/cuboid"
	"shapes-demo/shape/rect"
	"shapes-demo/shape/square"
)

func main() {
	r1 := &rect.Rect{L: 123.34, B: 123}
	var s1 square.Square = 25.23
	r2 := &rect.Rect{L: 1123.123, B: 123.123}
	var s2 square.Square = 125.23

	c1 := &cuboid.Cuboid{L: 123.34, B: 123, H: 12.34}
	// PrintAreaAndPerimeter(r1)
	// PrintAreaAndPerimeter(&s1)

	slice := make([]shape.IShape, 5)
	slice[0] = r1
	slice[1] = s1
	slice[2] = r2
	slice[3] = s2
	slice[4] = c1

	PrintIShape(slice)

}

func PrintAreaAndPerimeter(ishape shape.IShape) {
	typeOf := ishape.GetType()
	fmt.Println("Area of", typeOf, ":", ishape.Area())
	fmt.Println("Perimeter", typeOf, ":", ishape.Perimeter())
}

// Similar to dependency injection
func PrintIShape(slice []shape.IShape) {
	for _, ishape := range slice {
		if ishape != nil {
			typeOf := ishape.GetType()
			fmt.Println("Area of", typeOf, ":", ishape.Area())
			fmt.Println("Perimeter", typeOf, ":", ishape.Perimeter())
		}
	}
}

type Cuboid struct {
}

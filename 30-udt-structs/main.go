package main

import "fmt"

func main() {

	var r1 Rect
	r1.L = 123.123
	r1.B = 123.433
	fmt.Println(r1)

	a1 := Area(r1)
	p1 := Perimeter(r1)

	fmt.Println("Area of Rect r1:", a1)
	fmt.Println("Perimeter of Rect r1:", p1)
	fmt.Println("Area of Rect r1:", r1.A)
	fmt.Println("Perimeter of Rect r1:", r1.P)

	r2 := Rect{L: 123.123, B: 123.23}
	a2 := r2.Area()      // calling a methods on rect r2
	p2 := r2.Perimeter() // calling a methods on rect r2

	fmt.Println("Area of Rect r2:", a2)
	fmt.Println("Perimeter of Rect r2:", p2)

	//fmt.Println("Length and Bredth of Rect r2:", r2.L, r2.B)
	fmt.Println("Area of Rect(A field) r2:", r2.A)
	fmt.Println("Perimeter of Rect(P field) r2:", r2.P)

	// pointer receiver

	r3 := Rect{L: 123.123, B: 123.23}
	a3 := r3.AreaP()      // calling a methods on rect r2
	p3 := r3.PerimeterP() // calling a methods on rect r2

	// (&r3).AreaP() // very old way of calling

	fmt.Println("Area of Rect r3:", a3)
	fmt.Println("Perimeter of Rect r3:", p3)

	fmt.Println("Area of Rect(A field) r3:", r3.A)
	fmt.Println("Perimeter of Rect(P Field) r3:", r3.P)

	r4 := &Rect{L: 123.123, B: 123.23}
	a4 := r4.AreaP()      // calling a methods on rect r2
	p4 := r4.PerimeterP() // calling a methods on rect r2

	fmt.Println("Area of Rect r4:", a4)
	fmt.Println("Perimeter of Rect r4:", p4)

	fmt.Println("Area of Rect(A field) r4:", r4.A)
	fmt.Println("Perimeter of Rect(P Field) r4:", r4.P)

	r5 := new(Rect)

	r5.L = 123.23
	r5.B = 123.45
	a5 := r5.AreaP()
	p5 := r5.PerimeterP() // calling a methods on rect r2

	fmt.Println("Area of Rect r5:", a5)
	fmt.Println("Perimeter of Rect r5:", p5)

	fmt.Println("Area of Rect(A field) r5:", r5.A)
	fmt.Println("Perimeter of Rect(P Field) r5:", r5.P)
}

type Rect struct {
	A, P float32 // Area and Perimeter two fields
	L, B float32
}

// methods
// methods contain receivers
func (r Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

func (r *Rect) AreaP() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) PerimeterP() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

// normal functions
func Area(r Rect) float32 {
	r.A = r.L * r.B
	return r.A
}

func Perimeter(r Rect) float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

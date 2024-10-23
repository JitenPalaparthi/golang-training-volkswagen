package main

import "fmt"

func main() {

	s1 := Square{float32: 12.34}

	fmt.Println("Area of s1:", s1.Area())

	s2 := New(25.45)
	fmt.Println("Area of s2:", s2.Area())

	c1 := Cuboid{float32: 123.23, float2: 1231, myfloat: 123.34}
	fmt.Println(c1)

}

type Square struct {
	float32
}

type float2 = float32
type myfloat float32

type Cuboid struct {
	float32 // l
	float2  //b
	myfloat //h
}

// Kind of a constructor
// go's idamatic approach
func New(s float32) *Square {
	sq := new(Square)
	sq.float32 = s
	return sq
}

func (s *Square) Area() float32 {
	return s.float32 * s.float32
}

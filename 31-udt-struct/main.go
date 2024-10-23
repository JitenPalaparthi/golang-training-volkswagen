package main

import "fmt"

func main() {

	s1 := Square{float32: 12.34}

	fmt.Println("Area of s1:", s1.Area())

	s2 := New(25.45)
	fmt.Println("Area of s2:", s2.Area())

}

type Square struct {
	float32
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

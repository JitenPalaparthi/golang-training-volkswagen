package main

import (
	_ "custom-shapes/cuboid" //blank identifier
	"custom-shapes/rect"
	sq "custom-shapes/square"
	"log"
)

func main() {

	r1, err := rect.New(12.23, 12.43)
	if err != nil {
		log.Println(err.Error())
	}

	a1, p1 := r1.Area(), r1.Perimeter()
	log.Println("Area of Rect r1:", a1, "Perimeter of Rect r1:", p1)
	var s1 sq.Square = 25.25

	a2, p2 := s1.Area(), s1.Perimter()
	log.Println("Area of Square s1:", a2, "Perimeter of Square s1:", p2)

}

// to create a custom package , there should be a directory
// the name of the directory is generally(idiomatically) the name of the package

// go does not have access modifiers or specifies like public/pub private, protected etc..
// go has only exported or unexported types or fileds , functions, methods etc..
// or also called as restricted or not restricted
// any Identifier that stats with Uppercase it is exportable  (like public)
// any Identifier that stats with lowerCase it is not exportable (like private)

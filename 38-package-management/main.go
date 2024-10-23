package main

import (
	"custom-shapes/rect"
	"log"
)

func main() {

	r1, err := rect.New(12.23, 12.43)
	if err != nil {
		log.Println(err.Error())
	}

	a1, p1 := r1.Area(), r1.Perimeter()
	log.Println("Area of Rect r1:", a1, "Perimeter of Rect r1:", p1)

}

// to create a custom package , there should be a directory
// the name of the directory is generally(idiomatically) the name of the package

// go does not have access modifiers or specifies like public/pub private, protected etc..
// go has only exported or unexported types or fileds , functions, methods etc..
// or also called as restricted or not restricted
// any Identifier that stats with Uppercase it is exportable  (like public)
// any Identifier that stats with lowerCase it is not exportable (like private)

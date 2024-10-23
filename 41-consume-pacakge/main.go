package main

import (
	"fmt"

	"github.com/JitenPalaparthi/shapes/shape/rect"
)

func main() {

	r1, err := rect.New(10.23, 20.34)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Area:", r1.Area())
		fmt.Println("Perimeter:", r1.Perimeter())
	}
}

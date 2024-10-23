package main

import "fmt"

func main() {

	var a1 struct {
		Line1, City, Country string
		F1                   func()
	} // anonymous struct

	a1.City = "Hyd"
	a1.Country = "India"
	a1.Line1 = "Street1"
	a1.F1 = Greet

	fmt.Println(a1)
	a1.F1()

	// These two are also anonymous struct
	a2 := struct{ Name, Email string }{Name: "Jiten", Email: "Jitenp@Outlook.com"}
	a3 := struct{ Name, Email string }{"Jiten", "Jitenp@Outlook.com"}
	fmt.Println(a2)
	fmt.Println(a3)

}

func Greet() {
	fmt.Println("Hello World")
}

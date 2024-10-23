package main

import "fmt"

func main() {

	c1 := Company{Name: "ABCD", Email: "info@abcd.com", Contact: "12312312", Address: Address{Line1: "street1 chalakuzhy road", City: "Trivandrum", State: "Kerala", Country: "India", PinCode: "695411"}}

	fmt.Println(c1)

	fmt.Println("Address of c1", c1.Address)
	c1.Address.Print()

	a1 := c1.Address
	fmt.Println("Address of c1", a1)
	a1.Print()

}

type Company struct {
	Name    string
	Email   string
	Contact string
	Address Address // composition
}

type Address struct {
	Line1, City, State, Country, PinCode string
}

func (a *Address) Print() {
	fmt.Printf("Lin1:%v City:%v State:%v Country:%v PinCode:%v\n", a.Line1, a.City, a.State, a.Country, a.PinCode)
}

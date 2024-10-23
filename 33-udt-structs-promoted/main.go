package main

import "fmt"

func main() {

	c1 := Company{Name: "ABCD", Email: "info@abcd.com", Contact: "12312312", Status: "Active", Address: Address{Line1: "street1 chalakuzhy road", City: "Trivandrum", State: "Kerala", Country: "India", PinCode: "695411", Status: "Available"}}

	println("Country:", c1.Country)
	println("Company status", c1.Status)
	println("Address status", c1.Address.Status)

	c1.Print()

}

type Company struct {
	Name    string
	Email   string
	Contact string
	Status  string
	Address // promoted field. Just give the name of the struct
}

type Address struct {
	Line1, City, State, Country, PinCode string
	Status                               string
}

func (a *Address) Print() {
	fmt.Printf("Lin1:%v City:%v State:%v Country:%v PinCode:%v\n", a.Line1, a.City, a.State, a.Country, a.PinCode)
}

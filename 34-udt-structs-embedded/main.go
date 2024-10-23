package main

import "fmt"

func main() {

	c1 := Company{Name: "ABCD", Email: "info@abcd.com", Contact: "12312312", Address: struct{ Line1, City, State, Country, PinCode string }{Line1: "street1 chalakuzhy road", City: "Trivandrum", State: "Kerala", Country: "India", PinCode: "695411"}}

	fmt.Println(c1)

	fmt.Println("address of a company c1")
	c1.PrintAddress()

	//fmt.Println("Address of c1", c1.Address)

}

type Company struct {
	Name    string
	Email   string
	Contact string
	Address struct { // embedded struct
		Line1, City, State, Country, PinCode string
	}
}

func (a *Company) PrintAddress() {
	fmt.Printf("Lin1:%v City:%v State:%v Country:%v PinCode:%v\n", a.Address.Line1, a.Address.City, a.Address.State, a.Address.Country, a.Address.PinCode)
}

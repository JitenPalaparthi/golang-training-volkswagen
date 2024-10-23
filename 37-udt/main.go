package main

import (
	"errors"
	"fmt"
	"reflect"
)

type mymap map[string]any

func (m mymap) GetKeys() ([]string, error) {

	if m == nil {
		return nil, errors.New("nil mymap")
	}

	keys := make([]string, len(m))
	index := 0
	for k := range m {
		keys[index] = k
		index++
	}

	return keys, nil
}

func (m mymap) GetValues() (values []any, err error) {
	if m == nil {
		return nil, errors.New("nil mymap")
	}

	values = make([]any, len(m))
	index := 0
	for _, v := range m {
		values[index] = v
		index++
	}

	return values, nil
}

func (m mymap) Display() {
	for k, v := range m {
		switch v := v.(type) {
		case func():
			{
				v()
			}
		default:
			fmt.Println("Key:", k, "Value:", v, "Type:", reflect.TypeOf(v))
		}
	}
}

func Greet() {
	fmt.Println("Hello World!")
}
func main() {

	m1 := make(mymap)
	m1["1001"] = "Jiten"
	m1["1002"] = 7
	m1["1003"] = true
	m1["1004"] = "Done"
	m1["1005"] = Greet

	if keys, err := m1.GetKeys(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Keys:", keys)
	}

	if values, err := m1.GetValues(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Values:", values)
	}

	m1.Display()

	m2 := make(map[string]any) // not a user defined map it is standard map
	m2["1001"] = "Jiten"
	m2["1002"] = 7
	m2["1003"] = true
	m2["1004"] = "Done"
	m2["1005"] = Greet
	fmt.Println("m2 standard map after type casting and calling mymap methods")

	if keys, err := mymap(m2).GetKeys(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Keys:", keys)
	}

	if values, err := mymap(m2).GetValues(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Values:", values)
	}

	mymap(m2).Display() // type caste and call

	m3 := make(map[string]string)
	m3["1001"] = "Jiten"
	m3["1002"] = "7"
	m3["1003"] = "true"
	m3["1004"] = "Done"

	//mymap(m3).Display() // cant type caste

}

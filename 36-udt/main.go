package main

import (
	"fmt"
	"reflect"
)

type integer = int // This is not a new type , alias only

type myint int // This is a new type but created from existing type

type myint64 uint64

func (m myint) ToString() string {
	return fmt.Sprint(m)
}

func (m myint) Square() myint {
	return m * m
}

func main() {

	var num1 myint = 100
	str1 := num1.ToString()
	fmt.Println("num1:", num1, "Type:", reflect.TypeOf(num1))
	fmt.Println("str1 from num1:", str1, "Type:", reflect.TypeOf(str1))

	var num2 int = 300
	var float1 float32 = 123.23

	str2 := myint(num2).ToString()
	sq2 := myint(num2).Square()
	str3 := myint(float1).ToString()

	var num3 myint = myint(num2)
	var num4 integer // equallent to int in all aspects

	str4 := myint(num4).ToString()

	fmt.Println("str2 from num2:", str2, "Type:", reflect.TypeOf(str2))
	fmt.Println("str3 from float1:", str3, "Type:", reflect.TypeOf(str3))
	fmt.Println("str4 from num4:", str4, "Type:", reflect.TypeOf(str4))

	fmt.Println("sq2 from num2:", sq2)
	fmt.Println("casted num2 of int to num3 of myint", num3)

	var s1 Square = 123.123
	fmt.Println("Area of s1:", s1.Area())
	fmt.Println("Perimeter of s1:", s1.Perimeter())

	var num5 myint64 = 1231231231

	str5 := myint(num5).ToString()
	fmt.Println("str5 from num4:", str5, "Type:", reflect.TypeOf(str5))

}

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}

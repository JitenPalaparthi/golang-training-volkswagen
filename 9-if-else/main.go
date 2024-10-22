package main

import "fmt"

func main() {

	var num int = 1233
	var ok bool = num%2 == 0

	if ok {
		println(num, "is even")
	} else {
		println(num, "is odd")
	}

	// another one

	var (
		age    uint8
		gender rune
	)

	age = 18
	gender = 'F'

	// age >=18 && (false || true)
	// age>=18 && true
	// true && true
	// true

	ok1 := age >= 18 && (gender == 'f' || gender == 70)
	println(age >= 18 && (gender == 'f' || gender == 70))
	if ok1 /*gender == 102*/ {
		fmt.Println("She is eligible for marriage becase age is ", age)
	} else if age >= 21 && (gender == 'm' || gender == 'M') {
		fmt.Println("He is eligible for marriage becase age is ", age)
	} else {
		println("not eligible for marriage")
	}

}

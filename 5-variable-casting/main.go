package main

import "fmt"

type interger = int
type char = rune // alias for int32 and in all aspects both are same

func main() {

	var num1 interger = 12312
	var char1 rune = 'A'
	var char2 rune = '界'
	var char3 int32 = '你'
	var char4 char = char3
	var num2 uint8 = 65
	var num3 = 1312312321
	var byte1 byte = 65

	fmt.Println(num1, char1, char2, char3, char4)
	fmt.Println(string(char4))
	fmt.Println(string(num2))
	fmt.Println(string(num3))
	fmt.Println(string(byte1))
}

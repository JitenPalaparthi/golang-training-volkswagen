package main

import (
	"fmt"
	"reflect"
	"strconv"
	//_ "time"
)

func main() {

	str1 := "231312"

	/// var num1 int = int(str1) // type casting does not work here

	var num1 int
	//num1, err := strconv.Atoi(str1)
	num1, _ = strconv.Atoi(str1)
	fmt.Println("Value:", num1, "Type:", reflect.TypeOf(num1))
	fmt.Printf("Value: %d Type: %T\n", num1, num1)

	num2 := 121231221

	str2 := strconv.Itoa(num2)
	fmt.Printf("Value: %s Type: %T\n", str2, str2)

	str3 := fmt.Sprint(num2)
	fmt.Printf("Value: %s Type: %T\n", str3, str3)

	var (
		num3   = 12312
		float1 = 12312.123123
		ok1    = true
	)

	str4 := fmt.Sprintf("%d %f %v", num3, float1, ok1)

	fmt.Printf("Value: %s Type: %T\n", str4, str4)

	var num4 uint64 = 121231221

	str5 := strconv.Itoa(int(num4))
	fmt.Printf("Value: %s Type: %T\n", str5, str5)

	str6 := fmt.Sprint(num4)
	fmt.Printf("Value: %s Type: %T\n", str6, str6)

	// more on strconv

	str7 := "101010" // 42
	num5, err := strconv.ParseInt(str7, 2, 64)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(num5)
	}

	str8 := "2A" // 42
	num6, err := strconv.ParseInt(str8, 16, 64)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(num6)
	}
}

// nil there , but not null
// _ blank identifier

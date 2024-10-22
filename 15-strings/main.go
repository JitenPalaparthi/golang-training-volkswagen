package main

import (
	"fmt"
	"unsafe"
)

func main() {

	str1 := "Hello world!"

	str2 := "Hello 世界!"
	println(str1, "\n", str2)

	//reflect.TypeOf()

	b1 := []byte(str1)

	fmt.Println("address of str1 of byte array b1:", &b1[0])

	fmt.Println("Size of str1:", unsafe.Sizeof(str1), "address of str1", &str1)
	fmt.Println("Size of str2:", unsafe.Sizeof(str2), "address of str2", &str2)

	str1 = "Hello peaceful world!" // mutated the string

	fmt.Println("Size of str1:", unsafe.Sizeof(str1), "address of str1", &str1)

	b2 := []byte(str1)

	fmt.Println("address of str1 of byte array b2:", &b2[0])

}

/* string

str1
ptr: pointer to the initial charster // 8 bytes
len: the length of the string // 8 bytes
*/

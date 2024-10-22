package main

func main() {
	str1 := "Hello world!"
	str2 := "Hello 世界!"
	println(str1, "\n", str2)

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]), "  ")
	}
	println()

	for i := 0; i < len(str2); i++ {
		print(string(str2[i]), "  ")
	}
	println()

	for _, v := range str2 { // index and value
		//println("i:", i, "char:", string(v))
		print(string(v), " ")
	}
	println()

	for i := range str2 { // index and value
		//println("i:", i, "char:", string(v))
		print(i, " ")
	}
	println()
}

// range loop works for strings, arrays, slices, maps, channels

/* string

str1
ptr: pointer to the initial charster // 8 bytes
len: the length of the string // 8 bytes
*/

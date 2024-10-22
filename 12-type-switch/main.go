package main

func main() {

	var any1 any

	var num1 uint8 = 2
	any1 = num1

	switch any1 := any1.(type) { // type switch
	case uint8:
		println("square of uint8:", any1*any1)
	case uint16:
		println("square:", any1*any1)
	case uint32:
		println("square:", any1*any1)
	case uint64:
		println("square:", any1*any1)
	case int8:
		println("square:", any1*any1)
	case int16:
		println("square:", any1*any1)
	case int32:
		println("square:", any1*any1)
	case int64:
		println("square:", any1*any1)
	case int:
		println("square:", any1*any1)
	case uint:
		println("square:", any1*any1)
	case float32:
		println("square:", any1*any1)
	case float64:
		println("square:", any1*any1)

	default:
		println("cannot perform multiplication operations on this type")
	}
}

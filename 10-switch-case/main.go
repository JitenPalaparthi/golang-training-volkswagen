package main

func main() {

	var day uint8 = 41

	switch day {
	case 1:
		println("Sunday")
		println("happy day")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")

	default:
		println("noday")
	}

	var value uint = 120

	switch { // empty switch
	case value <= 50:
		println(value, "which is between 0-50")
	case value > 50 && value <= 100:
		println(value, "which is between 51-100")
	default:
		println(value, "which is greater than 100")
	}

	var char rune = 'e'

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U': // case with multiple options
		println(char, "is vovel")
	default:
		println(char, "is consonent or a special char")
	}

	// above logic in another way with empty switch
	switch {
	case char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U':
		println(char, "is vovel")
	default:
		println(char, "is consonent or a special char")
	}

}

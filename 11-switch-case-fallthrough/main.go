package main

func main() {
	num := 6
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough // This gives wrong result becase of this fallthrough , it falls to the body of the following case
	case num%12312 == 0: // this case is not checked as the above statement is satisfied and there is a fallthrough
		println(num, "is divisible by 12312")
	}
}

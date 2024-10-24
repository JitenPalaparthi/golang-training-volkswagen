package main

func main() {

	num := 1

loop:
	println(num)
	num++
	if num <= 10 {
		goto loop
	}

	println("out of loop")

}

package main

func main() {

	// 1-10

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	// for like while loop

	i := 1

	for i <= 10 {
		print(i, " ")
		i++
	}

	println()

	i = 1

	for {
		if i > 10 {
			break
		}
		print(i, " ")
		i++
	}
	println()
	println("Print even numbers")
	for i := 1; i <= 50; i++ {
		if i%2 == 1 {
			continue
		}
		print(i, " ")
	}
	println()

}

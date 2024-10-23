package main

func main() {

	f1, f2 := 0, 1

	fb := 0
	for fb <= 100 {
		fb = f1 + f2
		println(fb)
		f1, f2 = f2, fb

	}

}

// 1 1 2 3 5 8 13 21 34 55

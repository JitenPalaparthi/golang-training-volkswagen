package main

func main() {

	// ch := generate(100)
	// for v := range ch {
	// 	println(v)
	// }
	// for v := range generate(100) {
	// 	println(v)
	// }

	signal := make(chan struct{})

	go func() {
		for v := range generate(100) {
			println(v)
		}
		signal <- struct{}{} // signal , future pattern
	}()

	<-signal

}

func generate(num int) chan int { // generator pattern
	ch := make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

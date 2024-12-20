package main

func main() {

	ch := make(chan int)
	signal := make(chan struct{}) // empty struct channel

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i * i
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			println(<-ch)
		}
		signal <- struct{}{} // as soon as i receive all the data from ch .. send a signal to another channel
	}()

	<-signal // why it is blocked here?

	//runtime.Goexit()

	// s := struct{ i int }{i: 10}
	// empty := struct{}{} // no fields and no values , size is zero bytes
}

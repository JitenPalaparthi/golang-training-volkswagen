package main

func main() {

	ch := make(chan int, 10)      // when should and what number to be used as a buffer
	signal := make(chan struct{}) // empty struct channel

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i * i
		}
		close(ch)
	}()

	go func() {
		for c := range ch { // on a channel range receives the value until chan is closed
			println(c)
		}
		signal <- struct{}{} // as soon as i receive all the data from ch .. send a signal to another channel
	}()

	<-signal // why it is blocked here?

	//runtime.Goexit()

	// s := struct{ i int }{i: 10}
	// empty := struct{}{} // no fields and no values , size is zero bytes
}

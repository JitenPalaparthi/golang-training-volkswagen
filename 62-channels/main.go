package main

func main() {

	ch := make(chan int, 10) // when should and what number to be used as a buffer
	//signal := make(chan struct{}) // empty struct channel

	go sender(ch)
	// sig := receiverS(ch)
	// <-sig
	<-receiverS(ch)
	//go receiver(ch, signal)
	//<-signal // why it is blocked here?
}

func sender(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i * i
	}
	close(ch)
}
func receiver(ch <-chan int, signal chan<- struct{}) {
	for c := range ch { // on a channel range receives the value until chan is closed
		println(c)
	}
	signal <- struct{}{}
}

func receiverS(ch <-chan int) chan struct{} {
	signal := make(chan struct{}) // empty struct channel
	go func() {
		for c := range ch { // on a channel range receives the value until chan is closed
			println(c)
		}
		signal <- struct{}{}
		close(signal)
	}()
	return signal
}

package main

import "fmt"

func main() {
	var ch1 chan int     // channel that can hold int value, it is nil channel bcz not instantiated
	ch1 = make(chan int) // unbuffered channel
	go func() {
		v := <-ch1 // this is blocked for 3 seconds becase sending is not sending for 3 seconds
		println(v)
	}()
	ch1 <- 100 // sending 100 to a channel ch1
	fmt.Println("Done")
}

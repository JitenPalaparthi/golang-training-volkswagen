package main

import (
	"time"
)

func main() {
	var ch1 chan int     // channel that can hold int value, it is nil channel bcz not instantiated
	ch1 = make(chan int) // unbuffered channel
	//ch1 = make(chan int, 10) // buffered channel, can hold 10 ints at once.
	go func() {
		time.Sleep(time.Second * 3)
		//runtime.Goexit()
		// return
		// panic("something has gone wrong")
		ch1 <- 100 // sending 100 to a channel ch1
	}()
	v := <-ch1 // this is blocked for 3 seconds becase sending is not sending for 3 seconds
	println(v)
}

// 1. chan is a keyword to create a channel
// 2. make function is used to instantiate a channel
// 3. there are two kinds of channels, buffered and unbuffered.
// 		mostly we use unbuffered channels
// 4. channel can be nil , if not instantiated
// 5. to assign a value to a channel ch <- 100. To receive a value <- ch
// 6. IMP: in a goroutine a value is sent using a channel, the send is blocked until the
// receiver goroutine receives the value. The receiver is also blocked untile the sender sends the value.
// this is particulrly for unbuffered channels.
// for bufferec channels this is block is there only if the buffer if full

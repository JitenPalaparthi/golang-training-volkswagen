package main

import "time"

func main() {
	go println("Hello gorountines")
	go greet1() // this run as a goroutine
	go greet2("Hello VolksWagen folks!")
	println("I am dont with my execution")
	time.Sleep(time.Millisecond * 10) // this is not a solution
}

func greet1() {
	println("Hello world.. using goroutine")
}

func greet2(msg string) {
	println(msg)
}

// 1. main is also a goroutine
// 2. goroutines does not wait for other goroutines to finish off its execution
// 3. to create a goroutine a). it should be a function or a method b). use go keyword
// 4. in case of 2.w.r.t main, other goroutines may work , may not
// 5. goroutines may run in any order. Order of execution is not guaranteed.

// M  -> OS Threads
// P  -> Logical CPU , P number of threds are created by Go by default
// G  -> Go's goroutines

// Go uses an algorthm called work-stealing
// Global queue
// Local  queue

// ------- When a new goroutine is created
//

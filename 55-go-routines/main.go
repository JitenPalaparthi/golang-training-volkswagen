package main

import (
	"runtime"
	"sync"
)

var wg *sync.WaitGroup

func init() {
	wg = new(sync.WaitGroup)
}

func main() {
	println("Main has started")
	wg.Add(3)
	go func(wg *sync.WaitGroup, num int) {
		for i := 1; i <= num; i++ {
			if i%2 != 0 {
				println("Odd", i)
			}
		}
		wg.Done()
	}(wg, 20)
	//wg.Add(1)
	go func(wg *sync.WaitGroup) {
		printEven("even-func-1", 10)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		go primeNumber()
		wg.Done()
	}(wg)
	//wg.Add(1)
	// go func(wg *sync.WaitGroup) {
	// 	printEven("even-func-2", 400)
	// 	wg.Done()
	// }(wg)

	wg.Wait() // wait until the wg counter becomes zero(0)
	println("Main is going to be exited")
}

func printEven(fname string, num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			println(fname, ":Even", i)
		}
	}
}

func primeNumber() {
	i := 1
	for {
		c := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				c++
				if c >= 3 {
					break
				}
			}
		}
		if c <= 2 {
			println("prinme number:", i)
		}
		i++
		if i >= 100 {
			runtime.Goexit()
		}
	}
}

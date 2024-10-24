package main

import "sync"

var (
	counter int = 0
	wg      *sync.WaitGroup
)

func main() {
	wg = new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go increment(wg)
		}
		wg.Done()
	}(wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go decrement(wg)
		}
		wg.Done()
	}(wg)
	wg.Wait()
	println(counter)
}

func increment(wg *sync.WaitGroup) {
	counter++
	wg.Done()
}

func decrement(wg *sync.WaitGroup) {
	counter--
	wg.Done()
}

// do not communicate by sharing memory,intead share memory by communicating

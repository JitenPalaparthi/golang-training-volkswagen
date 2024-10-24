package main

import "sync"

// this is a solution for 56-mutex-problem-program
var (
	counter int = 0 // when a global variable or date is used/mutated by multiple goroutines use mutes.
	wg      *sync.WaitGroup
	mu      *sync.Mutex
)

func main() {
	wg = new(sync.WaitGroup)
	mu = new(sync.Mutex)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go increment(wg, mu)
		}
		wg.Done()
	}(wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go decrement(wg, mu)
		}
		wg.Done()
	}(wg)
	wg.Wait()
	println(counter)
}

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func decrement(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter--
	mu.Unlock()
	wg.Done()
}

// do not communicate by sharing memory,intead share memory by communicating

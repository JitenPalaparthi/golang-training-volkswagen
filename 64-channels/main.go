package main

import (
	"fmt"
	"time"
)

func main() {

	elapse := time.After(time.Second * 2)
	ch1 := generate("multiply by 2", 2)
	ch2 := generate("multiply by 3", 3)
	ch3 := make(chan int)
	ch4 := make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		close(ch4) // when the channel is closed, select case gets the notifacation
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 100)
			ch3 <- i * i
		}
		close(ch3)
	}()

outer:
	for {
		select {
		case v1 := <-ch1:
			println(v1)
		case v2 := <-ch2:
			println(v2)
		case v3, ok := <-ch3:
			if !ok {
				println("closed")
			} else {
				println("square of", v3)
			}
		case v4, ok := <-ch4:
			if !ok {
				println("simple  ch4--------------->:closed")
			} else {
				println("nothing comes here", v4)
			}
		case <-elapse:
			break outer
		}
	}
	fmt.Println("Process exit")

}

func generate(name string, n int) chan string { // generator pattern
	ch := make(chan string)
	go func() {
		i := 1
		for {
			time.Sleep(time.Millisecond * 500)
			ch <- fmt.Sprint(name, "-->", (i * n))
			i++
		}
		close(ch)
	}()
	return ch
}

// func generate(num int) chan int { // generator pattern
// 	ch := make(chan int)
// 	go func() {
// 		for i := 1; i <= num; i++ {
// 			ch <- i * i
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

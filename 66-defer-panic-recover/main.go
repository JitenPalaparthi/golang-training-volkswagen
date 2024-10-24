package main

import (
	"fmt"
	"time"
)

func main() {
	//defer recoverThis()
	go divideBy(100) // panic which is divide by 0
	go multiplyBy(100)
	go multiplyBy(100)
	time.Sleep(time.Second * 1)
}

func recoverThis() {
	if r := recover(); r != nil {
		fmt.Println("not doing anything, just recovering", r)
	}
}

// by default, panic panics the whole call stack.

func divideBy(num int) {
	defer println("divideBy function finished executing or crashed")
	defer recoverThis()
	for i := 10; i >= 0; i-- { // when there is a paninc ,the programe crashes but all defer functions are executed
		if i == 0 {
			panic("There seems to be divide by zero error. So panicing")
		}
		println("Divide by result:", num/i)
	}
}

func multiplyBy(num int) {
	defer println("multiplyBy function finished executing")
	defer recoverThis()
	for i := 10; i >= 0; i-- {
		println("Multiply", num*i)
	}
}

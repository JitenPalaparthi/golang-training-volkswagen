package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("task completed")
			case <-ctx.Done():
				fmt.Println("Task completed:", ctx.Err())
				runtime.Goexit()
			}
		}
	}()
	time.Sleep(time.Second * 2)
	cancel()

	time.Sleep(time.Second * 3)

}

// https://www.youtube.com/playlist?list=PLJE7PIP1qj_Rn9vq4V4jGJbj5KqEIWSUc

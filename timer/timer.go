package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	timeout := time.NewTimer(time.Second)
	defer timeout.Stop()

	select {
	case <-ch:
		fmt.Println("Work done")
	case <-timeout.C:
		fmt.Println("Work timeout")
	}
}

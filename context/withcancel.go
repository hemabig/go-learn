package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("work done")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("we do some work")
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()
}

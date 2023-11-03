package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("work canel")
				return
			default:
				fmt.Println("we do some work")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	//time.Sleep(8 * time.Second)
}

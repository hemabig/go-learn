package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go func(ctx context.Context) {
		deadline, ok := ctx.Deadline()
		if ok {
			fmt.Println("cancle at:", deadline)
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
		}
	}(ctx)

	time.Sleep(8 * time.Second)
}

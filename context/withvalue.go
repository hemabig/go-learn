package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "name", "Joe")

	go func(ctx context.Context) {
		value := ctx.Value("name")
		fmt.Println("value:", value)
	}(ctx)

	time.Sleep(1 * time.Second)
}

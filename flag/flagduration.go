package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	sleep time.Duration
)

func init() {
	flag.DurationVar(&sleep, "period", 1*time.Second, "sleep period")
}

func main() {
	flag.Parse()
	fmt.Println("Duration flag: ", sleep)
}

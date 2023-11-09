package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

type Interval []time.Duration

func (i *Interval) String() string {
	return fmt.Sprint(*i)
}

func (i *Interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}

	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

var (
	intervalflag Interval
)

func init() {
	flag.Var(&intervalflag, "detaT", "comma-seperated list of intervals to use between events")
}

func main() {
	flag.Parse()
	fmt.Println(intervalflag)
}

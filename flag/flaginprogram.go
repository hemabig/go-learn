package main

import (
	"flag"
	"fmt"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func main() {
	args := []string{"--intflag", "12", "--boolflag", "true"}
	parse := flag.NewFlagSet("MyFlagSet", flag.ContinueOnError)
	parse.IntVar(&intflag, "intflag", 0, "int flag value")
	parse.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	parse.StringVar(&stringflag, "stringflag", "default", "string flag value")
	parse.Parse(args)

	fmt.Println("int flag: ", intflag)
	fmt.Println("bool flag: ", boolflag)
	fmt.Println("string flag: ", stringflag)

	//没有解析的参数 -- 后面的参数
	fmt.Println(flag.Args())
	fmt.Println("Non-Flag Argument Count:", flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("Argument %d: %s\n", i, flag.Arg(i))
	}

	//解析的参数的个数
	fmt.Println("Flag Count:", flag.NFlag())
}

package main

//char *cs = "Hello World";
import "C"

import (
	"diffpackage/cgo_helper"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(C.cs))
	//fmt.Println(reflect.Type(cgohelper.CChar))
	cgo_helper.SayHello(C.cs)
}

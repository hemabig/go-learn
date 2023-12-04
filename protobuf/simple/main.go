package main

import (
	"fmt"
	message "protobuf_test/message"

	"github.com/golang/protobuf/proto"
)

func main() {
	p := &message.Request{Name: "Hello World", Age: 12, Hobbies: []string{"Jump", "Run"}}
	data, error := proto.Marshal(p)
	if error != nil {
		fmt.Println("Marshal error")
		return
	}

	fmt.Println(data)

	p1 := &message.Request{}

	proto.Unmarshal(data, p1)

	fmt.Println(p1.GetName(), p1.GetAge(), p1.GetHobbies())
}

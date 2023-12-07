package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Gender  string   `json:"gender"`
	Hobbies []string `json:"Hobbies"`
}

func main() {
	p := &student{
		Name:    "John",
		Age:     20,
		Gender:  "man",
		Hobbies: []string{"Jump", "Run", "Sleep"},
	}

	jsonbyte, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	jsonstring := string(jsonbyte)
	fmt.Println(jsonstring)

	var s student
	json.Unmarshal([]byte(jsonstring), &s)

	fmt.Printf("%v\n", s)

	var data map[string]interface{}

	json.Unmarshal([]byte(jsonstring), &data)
	fmt.Println("name: ", data["name"])
	fmt.Println("Hobbies: ", data["Hobbies"])

}

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
}

func main() {
	p := Person{Name: "John", Age: 30}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	input := ""
	var person map[string]string = make(map[string]string)
	fmt.Print("Enter your name: ")
	_, _ = fmt.Scan(&input)
	person["name"] = input

	fmt.Print("Enter your address without space: ")
	_, _ = fmt.Scan(&input)
	person["address"] = input

	op, _ := json.Marshal(person)
	fmt.Println(string(op))

}

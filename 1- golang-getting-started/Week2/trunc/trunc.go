package main

import (
	"fmt"
)

func main() {
	var input float64
	fmt.Printf("Enter a float: ")
	_, _ = fmt.Scan(&input)

	output := int64(input)
	fmt.Print(output)
}

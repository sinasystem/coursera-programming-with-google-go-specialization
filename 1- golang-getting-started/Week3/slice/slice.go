package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	sli := make([]int, 0, 3)
	var input string
	for {
		fmt.Println("Enter ints to add, Enter x to exit")
		_, _ = fmt.Scan(&input)

		if strings.ToLower(input) == "x" {
			break
		}

		temp, _ := strconv.Atoi(input)
		sli = append(sli, temp)
		sort.Ints(sli)
		fmt.Println(sli)
	}

}

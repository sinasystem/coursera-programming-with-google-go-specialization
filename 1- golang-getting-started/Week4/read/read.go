package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type name struct {
	fname string
	lname string
}

func main() {
	names := make([]name, 0, 3)
	fmt.Println("Enter the name of the text file:")
	var fileName string
	fmt.Scan(&fileName)

	if !strings.HasSuffix(fileName, ".txt") {
		fileName = fmt.Sprint(fileName, ".txt")
	}

	fileName = strings.Trim(fileName, "\n")

	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var nameTemp name
		nameTemp.fname, nameTemp.lname = s[0], s[1]

		names = append(names, nameTemp)
	}

	fmt.Print(names)

}

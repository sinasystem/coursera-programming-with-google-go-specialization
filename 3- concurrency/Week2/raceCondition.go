package main

import (
	"fmt"
	"time"
)

func f1(count int) {
	for {
		count++
	}
}
func f2(count int) {
	for {
		fmt.Println(count)
		time.Sleep(time.Second)
	}
}

/*
This program creates two goroutines which access and modify the same variable concurrently.
The first goroutine is an infinite loop, incrementing the variable.
Second goroutine is another infinite loop which prints the variable and goes to sleep for 1 second.
Race condition occures because goroutines are not synchronized, and might execute concurrently.
The order of execution of routines are not guaranteed, so the printed value might not be the actual variable's value at that time.
*/
func main() {
	count := 0
	go f1(count)
	go f2(count)

	time.Sleep(5 * time.Second)
}

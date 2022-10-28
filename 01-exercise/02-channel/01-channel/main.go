package main

import (
	"fmt"
	"strconv"
)

func main() {
	channel := make(chan string)
	go func(a, b int) {
		c := a + b
		channel <- strconv.Itoa(c)
	}(1, 2)
	// TODO: get the value computed from goroutine
	c := <-channel
	fmt.Printf("computed value %v\n", c)
}

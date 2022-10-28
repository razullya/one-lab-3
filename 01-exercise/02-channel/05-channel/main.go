package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//TODO: create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.

	consumer := func(ch <-chan int) {
		// read values from channel

		for v := range ch {
			fmt.Printf("Received: %d\n", v)

		}
		fmt.Println("Done receiving!")
		wg.Done()
	}

	ch := owner(5)
	go consumer(ch)
}

func owner(n int) chan int {
	var ch chan int
	wg.Add(1)
	ch <- n
	return ch
}

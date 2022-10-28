package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)

		// TODO: send all iterator values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
			time.Sleep(500*time.Millisecond)
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}

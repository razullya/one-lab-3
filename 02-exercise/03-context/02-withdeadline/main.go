package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			// Report result.
			ch <- data{"123"}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute()
	d := <-ch
	fmt.Printf("work complete: %s\n", d)

}

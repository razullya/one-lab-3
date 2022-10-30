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

	compute := func() <-chan data {
		fmt.Println("deadline starts")
		ctx, _ := context.WithDeadline(context.Background(), <-time.After(time.Second*10))
		ch := make(chan data)
		go func() {
			defer close(ch)
			// Simulate work.
			fmt.Println("work")

			// Report result.
			ch <- data{"123"}
		}()
		<-ctx.Done()
		fmt.Println("deadline ends")
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute()
	d := <-ch
	fmt.Printf("work complete: %s\n", d)

}

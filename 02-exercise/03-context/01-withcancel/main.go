package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())

	generator := func(ctx context.Context) <-chan int {
		out := make(chan int)
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		go func() {
			for i := 0; ; i++ {
				if i == 5 {
					cancel()
					close(out)
					break
				}
				out <- r1.Intn(100)
			}
		}()
		return out
	}

	out := generator(ctx)
	for n := range out {
		fmt.Println(n)
	}
}

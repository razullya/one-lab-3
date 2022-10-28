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

	ctx, cancel := context.WithCancel(context.Background())
	
	k:=0
	generator := func(ctx context.Context) <-chan int {

		out := make(chan int)
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		go func() {
			for i := 0; i < 5; i++ {
				out <- r1.Intn(100)
			}
			close(out)
		}()
		return out
	}
	out := generator(ctx)
	fmt.Println(<-out)
	k++
	
	// Create a context that is cancellable.
	if k==5{
		cancel()
	}

}

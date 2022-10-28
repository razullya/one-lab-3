package main

import "fmt"

func main() {
	channel := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			channel <- i
		}
		close(channel)
	}()
	// TODO: range over channel to recv values
	for v := range channel {
		fmt.Println(v)

	}

}

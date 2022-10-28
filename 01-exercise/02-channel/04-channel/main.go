package main

import (
	"fmt"
	"sync"
)

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch chan string, msg string) {
	// send message on ch1
	ch <- msg
}

var m sync.Mutex

func relayMsg(ch1 chan string, ch2 chan string) {

	// recv message on ch1
	// send it on ch2

	ch2 <- <-ch1
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1, "Hello")
	go relayMsg(ch1, ch2)
	// recv message on ch2
	fmt.Println(<-ch2)
}

package main

import (
	"context"
	"fmt"
)

type database map[string]bool

type ctxKey string

var db database = database{
	"jane": true,
}

var key ctxKey = "key"

func main() {
	processRequest("jane")
}

func processRequest(userid string) {
	// TODO: send userID information to checkMemberShip through context for
	ctx := context.WithValue(context.Background(), key, userid)
	// map lookup.
	ch := checkMemberShip(ctx)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShip - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do map lookup
// sends the result on the returned channel.
func checkMemberShip(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do some database lookup
		userId := ctx.Value(key).(string)
		status := db[userId]
		ch <- status
	}()
	return ch
}

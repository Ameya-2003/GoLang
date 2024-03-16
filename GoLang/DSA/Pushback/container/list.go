package main

import (
	"container/list"
	"fmt"
	"time"
)

func main() {
	var intlist list.List
	intlist.PushBack("chotte")
	intlist.PushBack("jhaatu")
	intlist.PushBack("hai")

	for element := intlist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	//Buffered Channel

	// Create a buffered channel with capacity 3
	q := make(chan int, 3)

	// Send some values to the channel
	q <- 1
	q <- 2
	q <- 3

	// Receive values from the channel in FIFO order
	fmt.Println(<-q) // 1
	fmt.Println(<-q) // 2
	fmt.Println(<-q) // 3

	// Close the channel so that any  subsequent send will block until there is space available
	close(q)

	// Now we can safely receive from it again:
	fmt.Println(<-q) // Blocks because the channel is closed and empty

	// A select statement allows multiple communication operations to happen at once
	// Here's an example where we send on two channels and then receive from them simultaneously
	c1 := make(chan string)
	c2 := make(chan string)

	go func() { c1 <- "channel 1" }()
	go func() { time.Sleep(100 * time.Millisecond); c2 <- "channel 2" }()

	select {
	case x := <-c1:
		fmt.Printf("received %q\n", x)
	case x := <-c2:
		fmt.Printf("received %q\n", x)
	}

	// In this case, since both `c1` and `c2` have values ready to be received immediately, one of the cases gets executed randomly.
	// In this case, since both `c1` and `c2` have values ready for us, one of the cases gets executed
	// The choice between which case runs first is not guaranteed (it depends on the scheduler).
	// So if you want deterministic behavior, use a default case or explicitly list all possible cases.

}

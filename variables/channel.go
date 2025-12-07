package main

import "fmt"

/* How They Work:
Creation: ch := make(chan int) creates an unbuffered channel for integers.
Sending: ch <- 42 sends the integer 42 into the channel ch.
Receiving: val := <-ch receives a value from ch and stores it in val.
Synchronization: If a goroutine tries to receive from an empty channel, it waits. If a goroutine tries to send to a full channel (in buffered channels), it waits. 
*/

func main() {
	fmt.Println("channel:")
	u := make(chan int) // unbuffered channel of integers
	i := make(chan float64, 5) // buffered channel of float64 with capacity 5
	fmt.Println("u=", u)
	fmt.Println("i=", i)
}
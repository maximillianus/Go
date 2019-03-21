/*
	Channels are the pipes that connect concurrent goroutines.
	You can send values into channels from one goroutine and
	receive those values into another goroutine.
*/

package main

import "fmt"

func main() {
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// send value to channel
	go func() { messages <- "1" }()

	msg := <-messages
	fmt.Println(msg)
}

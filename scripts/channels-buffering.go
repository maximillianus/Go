/*
	Make channels without receiver.
	This is called buffering. And it has to have a buffer limit.
*/

package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
	specify if a channel only meant to send/receive
*/

package main

import "fmt"

func ping(pings chan<- string, msg string) {

	fmt.Println(msg, "into ping")
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// pass message in ping to msg var
	msg := <-pings
	fmt.Println(msg, "from ping")

	// pass msg to pongs channel
	fmt.Println(msg, "into pong")
	pongs <- msg

}

func main() {
	chn1 := make(chan string, 1)
	chn2 := make(chan string, 1)
	// pass message to chn1
	ping(chn1, "something")

	// pass msg from chn1 to chn2
	pong(chn1, chn2)

	// get msg out from chn 2
	fmt.Println(<-chn2, "out of pong")

}

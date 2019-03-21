/*
	select is useful to wait for 2 channel concurrently
*/

package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		// this program will finish in 2 seconds
		// since it follows the time of the longest execution
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// select can be used to wait values from chn1 and chn2
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

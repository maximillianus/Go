package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c2 <- "Every 1s"
			time.Sleep(1e3 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}

	// fmt.Scanf()

	// msg := <-c
}

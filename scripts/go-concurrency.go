package main

import (
	"fmt"
	// "sync"
	"time"
)

func main() {
	c := make(chan string)
	// var wg sync.WaitGroup
	// wg.Add(1)
	fmt.Println("Start")
	go func() {
		count("sheep", c)
		// wg.Done()
	}()
	for {
		msg, open := <-c
		// fmt.Println(open)
		if !open {
			break
		}
		fmt.Println(msg)

	}

	// go count("bird")
	// wg.Wait()
	// fmt.Println("End")
}

func count(thing string, c chan string) {
	for i := 1; i < 5; i++ {
		// output to terminal
		// fmt.Println("counting:", i, thing)
		// pass it to channel
		c <- thing
		time.Sleep(120 * time.Millisecond)
	}
	close(c)
}

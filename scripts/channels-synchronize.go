/*
	We can use channels to synchronize execution across goroutines.
	Here’s an example of using a blocking receive to wait for a
	goroutine to finish.
*/

package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// pass true to channel done
	done <- true
}

func main() {
	// init channel called 'done'
	done := make(chan bool, 1)

	// run channel in function
	go worker(done)

	<-done
}

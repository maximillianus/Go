package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// goroutine invoked concurrently
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// This will be executed first since this is not concurrent
	// and it is considered blocking.
	func(msg string) {
		fmt.Println(msg)
	}("blocking")

	/*
		Scanln waits for inputs from user.
		Without this function, we cannot see goroutine because the script
		terminates too fast.
	*/
	fmt.Scanln()
	fmt.Println("done")
}

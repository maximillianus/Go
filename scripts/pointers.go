package main

import "fmt"

func changeValue(iVal int) {
	iVal = 0
}

// Pay attention to the argument.
// There is `*` in front of int to show that
// we are taking integer pointer
func changePointer(iPtr *int) {
	// adding star in front to show that we are changing
	// value at the referenced memory address
	*iPtr = 0

	// What is the content o f *iPtr?
	fmt.Println("Content of *iPtr:", *iPtr)
}

func main() {
	i := 1

	fmt.Println("initial value of i:", i)

	// This is passing by value method
	changeValue(i)
	fmt.Println("i after value changed by value pass:", i)

	// This is passing by reference method or called  'dereferencing'
	// pass the memory address of i using `&`
	changePointer(&i)
	fmt.Println("i after value changed by value pass:", i)

	// check value of &i
	fmt.Println("What is value of &i:", &i)

	// pointing a variable to address of i, making it into a pointer var
	p := &i
	fmt.Println("What is value of p:", p)
	fmt.Println("What is value of *p:", *p)
	fmt.Println("What is value of &p:", &p)

}

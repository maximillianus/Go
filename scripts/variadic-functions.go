/*
Variadic functions is function that can be called with
multiple trailing arguments

Useful if we cant determine the number of arguments
*/

package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums)

	//sum them up
	var total int
	total = 0
	for _, num := range nums {
		total += num
	}

	return total
}

func avg(nums ...int) float32 {

	fmt.Println(nums)

	//sum them
	total := 0
	for _, num := range nums {
		total += num
	}

	// divide by len
	arrayLen := len(nums)

	var average float32
	average = float32(total) / float32(arrayLen)
	return average

}

func main() {
	x := avg(3, 8, 1, 2)
	// fmt.Println("Sum:", x)
	fmt.Println("Avg:", x)

	y := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println("Using func(slices...) method:", avg(y...))

}

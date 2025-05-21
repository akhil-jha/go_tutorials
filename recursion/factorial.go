package main

import (
	"fmt"
)

func calculateFactorial(number uint64) uint64 {
	if number == 0 {
		return 1
	}
	return number * calculateFactorial(number-1)

}

func main() {
	var myNumber uint64
	myNumber = 65 // this is the max limit
	fmt.Printf("Factorial of %d is %d\n", myNumber, calculateFactorial(myNumber))

}

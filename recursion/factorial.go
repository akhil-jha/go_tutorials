package main

import (
	"fmt"
)

func calculateFactorial(number int) int {
	if number > 0 {
		return number * calculateFactorial(number-1)
	} else {
		return 1
	}

}

func main() {
	myNumber := 100
	fmt.Printf("Factorial of %d is %d", myNumber, calculateFactorial(myNumber))

}

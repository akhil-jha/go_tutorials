package main

import "fmt"

func isEven(input_number int) bool {
	return input_number%2 == 0
}

func main() {
	input_number := 11
	fmt.Printf("The given number is %v", isEven(input_number))
}

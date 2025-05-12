// Find the Largest Number
// Accept three integers from the user and print the largest among them.

package main

import "fmt"

func largestNumber(number1 int, number2 int, number3 int) int {
	max_number := number1

	if number2 > max_number {
		max_number = number2
	} else if number3 > max_number {
		max_number = number3
	}
	return max_number

	// if number1 > number2 && number1 > number3 {
	// 	return number1
	// } else if number2 > number1 && number2 > number3 {
	// 	return number2
	// } else {
	// 	return number3
	// }
}

func main() {
	var number1, number2, number3 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&number1)
	fmt.Print("Enter Second number: ")
	fmt.Scan(&number2)
	fmt.Print("Enter Thrid number: ")
	fmt.Scan(&number3)

	fmt.Printf("Largest number: %v", largestNumber(number1, number2, number3))
}

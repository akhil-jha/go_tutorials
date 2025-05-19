package main

import "fmt"

func main() {
	myStringArray := [4]string{"India", "USA", "Japan", "UK"}

	for index, value := range myStringArray {
		fmt.Printf("%v: %v\n", index, value)
	}

	myIntAraay := [...]int{10, 11, 12, 13, 14, 15}

	for index, value := range myIntAraay {
		fmt.Printf("%v: %v\n", index, value)
	}
}

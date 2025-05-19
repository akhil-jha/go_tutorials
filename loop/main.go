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

	my2DArray := [3][4]string{{"India", "USA", "Japan", "UK"}, {"New Delhi", "Washington", "Tokyo", "London"}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%s", my2DArray[i][j])
		}
		fmt.Println("")
	}

	fmt.Println(my2DArray)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var number = 5
	// fmt.Println("Enter a number: ")
	// fmt.Scan(&number)

	/*
			------------------------------------------------------------------------------------------
			*
			* *
			* * *
			* * * *

			for i := 0; i < number; i++ {
				for j := 0; j < i; j++ {
					fmt.Printf("*")
				}
				fmt.Print("\n")
			}
			------------------------------------------------------------------------------------------
			* * * * *
			* * * *
			* * *
			* *
			*

			for i := 0; i < number; i++ {
					for j := 0; j < (number - i); j++ {
							print("*")
					}
					print("\n")
			}
			------------------------------------------------------------------------------------------
			****
			*  *
			*  *
			****

			for i := 0; i < number; i++ {
					for j := 0; j < number; j++ {
							if i == 0 || i == number-1 || j == 0 || j == number-1 {
									print("*")
							} else {
									print(" ")
							}
					}
					print("\n")
			}
			------------------------------------------------------------------------------------------
			****
			 ****
			  ****
			   ****

			for i:=0 ; i < number ; i++ {
				star := number
				spaces := i
				fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("*", stars))
			}
			------------------------------------------------------------------------------------------

			   *
			  ***
			 *****
			*******

			for i := 0; i < number; i++ {
				stars := 2*i + 1
				spaces := number - i - 1
				fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("*", stars))
		}
			------------------------------------------------------------------------------------------
			 ***********
			  *       *
			   *     *
		        *   *
		         * *
		          *
	*/

	for i := 0; i < number; i++ {
		stars := 1
		spaces := 1
		fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("*", stars))
	}
}

package main

func main() {
	var number = 4
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

			for i := 0; i < number; i++ {
				fmt.Printf("%s", strings.Repeat(" ", i))
				for j := 0; j < number; j++ {
					print("*")
				}
				fmt.Printf("\n")
			}
	*/
}

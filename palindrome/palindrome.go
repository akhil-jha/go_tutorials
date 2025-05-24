/*
Given an integer x, return true if x is a

, and false otherwise.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/
package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x > 0 {
		digitsArray := []int{}
		i := 0
		for x > 0 {
			digitsArray = append(digitsArray, x%10)
			x = x / 10
			i++
		}

		for i := 0; i < len(digitsArray)/2; i++ {
			if digitsArray[i] != digitsArray[len(digitsArray)-i-1] {
				return false
			}
		}
	}
	return true
}

func main() {
	number := 121212121
	fmt.Println(isPalindrome(number))
}

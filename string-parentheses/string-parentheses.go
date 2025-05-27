/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.



Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true



Constraints:

    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.
*/

package main

import "fmt"

func isValid(s string) bool {
	round := map[rune]int{'(': 0, ')': 0}
	square := map[rune]int{'[': 0, ']': 0}
	curly := map[rune]int{'{': 0, '}': 0}
	for _, value := range s {
		if value == '(' {
			round[value] += 1
		}
		if value == ')' {
			round[value] += 1
		}
		if value == '[' {
			square[value] += 1
		}
		if value == ']' {
			square[value] += 1
		}
		if value == '{' {
			curly[value] += 1
		}
		if value == '}' {
			curly[value] += 1
		}
	}
	if (round['(']+round[')']+square[']']+square[']']+curly['{']+curly['}'])%2 != 0 {
		return false
	}
	return true

}

func main() {
	myString := "()[]}{"
	fmt.Println(isValid(myString))
}

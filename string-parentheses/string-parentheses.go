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

type Stack struct {
	item []rune
}

func (s *Stack) Push(i rune) {
	s.item = append(s.item, i)
}

func (s *Stack) Pop() {
	s.item = s.item[:len(s.item)-1]
}

func (s *Stack) Seek() rune {
	return s.item[len(s.item)-1]
}

func isValid(s string) bool {
	stack := Stack{}

	for _, value := range s {
		if value == '(' || value == '{' || value == '[' {
			stack.Push(value)
		}
		if value == ')' {
			if len(stack.item) == 0 || stack.Seek() != '(' {
				return false
			}
			stack.Pop()
		}
		if value == '}' {
			if len(stack.item) == 0 || stack.Seek() != '{' {
				return false
			}
			stack.Pop()
		}
		if value == ']' {
			if len(stack.item) == 0 || stack.Seek() != '[' {
				return false
			}
			stack.Pop()
		}

	}
	return len(stack.item) == 0
}

func main() {
	myString := "()[]}{"
	fmt.Println(isValid(myString))
}

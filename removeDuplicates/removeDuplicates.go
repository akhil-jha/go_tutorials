/*
You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.
We repeatedly make duplicate removals on s until we no longer can.
Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.

Example 1:

Input: s = "abbaca"
Output: "ca"
Explanation:
For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".

Example 2:
Input: s = "azxxzy"
Output: "ay"

Constraints:

	1 <= s.length <= 105
	s consists of lowercase English letters.
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

func removeDuplicates(s string) string {
	seen := Stack{}
	for _, value := range s {
		if len(seen.item) > 0 && seen.Seek() == value {
			seen.Pop()
		} else {
			seen.Push(value)
		}

	}
	return string(seen.item[:])
}

func main() {
	s := "abbacca"
	fmt.Println(removeDuplicates(s))
}

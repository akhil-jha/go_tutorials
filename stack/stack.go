package main

import "fmt"

type Stack struct {
	item []int
}

func (s *Stack) Push(num int) {
	s.item = append(s.item, num)
}

func (s *Stack) Pop() int {
	popped := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]
	return popped
}

func main() {
	stack := Stack{}

	stack.Push(20)
	stack.Push(30)
	stack.Push(90)

	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)

}

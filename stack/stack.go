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

func (s *Stack) Seek() int {
	return s.item[len(s.item)-1]
}

func main() {
	stack := Stack{}

	fmt.Println(stack)

	stack.Push(20)
	stack.Push(30)
	stack.Push(90)

	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)

	fmt.Println(stack.Seek())

}

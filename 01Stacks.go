package main

import "fmt"

func main() {
	fmt.Println("Hello")
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	removed := myStack.Pop()
	fmt.Println(removed)
	fmt.Println(myStack)
}

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]

	s.items = s.items[:l]
	return toRemove
}

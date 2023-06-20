package main

import (
	"fmt"
)

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func main() {
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	s.Push(30)
	v, ok := s.Pop()
	fmt.Println(v, ok)
	v, ok = s.Pop()
	fmt.Println(v, ok)
	v, ok = s.Pop()
	fmt.Println(v, ok)
	v, ok = s.Pop()
	fmt.Println(v, ok)
}

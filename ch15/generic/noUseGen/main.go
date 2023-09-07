package main

import "fmt"

func main() {
	noGen()
}

func noGen() {
	var s stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push("hello")
	var top, _ = s.Pop()
	fmt.Println(s)
	fmt.Println(top)

}

type stack struct {
	vals []interface{}
}

func (s *stack) Push(val interface{}) {
	s.vals = append(s.vals, val)
}

func (s *stack) Pop() (interface{}, bool) {
	if len(s.vals) == 0 {
		return nil, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:(len(s.vals) - 1)]
	return top, true
}

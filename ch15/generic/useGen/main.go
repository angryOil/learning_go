package main

import "fmt"

func main() {
	useGen()
}

func useGen() {
	var s stack[string]
	s.Push("1")
	s.Push("2")
	s.Push("3")
	s.Push("hello")
	s.Push("world")
	var top, ok = s.Pop()
	fmt.Println(top, ok)
	fmt.Println(s)
	fmt.Println(s.Contains("3"))
	fmt.Println(s.Contains("34"))
}

// 비교를 지원하려면

//type stack[T any] struct {
//	vals []T
//}

//아래로 변경

type stack[T comparable] struct {
	vals []T
}

func (s *stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

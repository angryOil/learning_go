package main

import "fmt"

func useNilMethod() {
	var it *IntTree
	it = it.insert(5)
	it = it.insert(3)
	it = it.insert(10)
	it = it.insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.insert(val)
	} else if val > it.val {
		it.right = it.right.insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

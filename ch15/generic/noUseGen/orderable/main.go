package main

import (
	"fmt"
	"strings"
)

func main() {
	useOrderable()
}

func useOrderable() {
	var it *tree[orderableInt]
	it = it.insert(10)
	it = it.insert(11)
	it = it.insert(20)
	it = it.insert(15)
	it = it.insert(orderableInt(22))
	fmt.Println(it)
}

type orderable[T any] interface {
	Order(T) int
}

type orderableInt int

func (oi orderableInt) Order(val orderableInt) int {
	return int(oi - val)
}

type orderString string

func (os orderString) Order(val orderString) int {
	return strings.Compare(string(os), (string)(val))
}

type tree[T orderable[T]] struct {
	val         T
	left, right *tree[T]
}

func (t *tree[T]) insert(val T) *tree[T] {
	if t == nil {
		return &tree[T]{val: val}
	}
	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.insert(val)
	case comp > 0:
		t.right = t.right.insert(val)
	}
	return t
}

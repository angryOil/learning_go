package main

import (
	"fmt"
	"strings"
)

func main() {
	var it *tree
	it = it.insert(orderableInt(3))
	it = it.insert(orderableInt(5))
	it = it.insert(orderableInt(6))
	it = it.insert(orderableInt(7))
	it = it.insert(orderableInt(1))
	it = it.insert(orderString("as"))
	fmt.Println(it.right.right.right)

	//var sTree *tree
	//sTree = sTree.insert(orderString("hello"))
	//sTree = sTree.insert(orderString("world"))
	//sTree = sTree.insert(orderString("joy"))
	//
	//fmt.Println(sTree)
}

type orderable interface {
	Order(interface{}) int
}

type tree struct {
	val         orderable
	left, right *tree
}

func (t *tree) insert(val orderable) *tree {
	if t == nil {
		return &tree{val: val}
	}
	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.insert(val)
	case comp > 0:
		t.right = t.right.insert(val)
	}
	return t
}

type orderableInt int

func (oi orderableInt) Order(val interface{}) int {
	return int(oi - val.(orderableInt))
}

type orderString string

func (os orderString) Order(val interface{}) int {
	return strings.Compare(string(os), val.(string))
}

type stack[T any] struct {
}

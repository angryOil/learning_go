package main

import "fmt"

func callByValue() {
	p := per{}
	i := 2
	s := "hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}

type per struct {
	age  int
	name string
}

func modifyFails(i int, s string, p per) {
	i = i * 2
	s = "goodBye"
	p.name = "bob"
}

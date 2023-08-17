package main

import "fmt"

func first() {
	fmt.Println(divs(11, 3))
	namedParameter(myFuncOpts{
		firstName: "hello",
		lastName:  "world",
	})
}

func divs(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

type myFuncOpts struct {
	firstName string
	lastName  string
	age       int
}

func namedParameter(opts myFuncOpts) {
	fmt.Println(opts)
}

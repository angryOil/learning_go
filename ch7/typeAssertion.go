package main

import "fmt"

func typeAssertion() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	//i2 := i.(int)
	//i2 := i.(MyInt)
	i2, ok := i.(int)
	if !ok {
		fmt.Println("unexpected type for ", i)
	} else {
		fmt.Println(i2 + 1)
	}
}

type MyInt int

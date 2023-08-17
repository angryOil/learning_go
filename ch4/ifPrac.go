package main

import (
	"fmt"
	"math/rand"
)

func ifPrac() {
	firIfPrac()
	secIfPrac()
}

func firIfPrac() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("wow 0")
	} else if n >= 5 {
		fmt.Println("big")
	} else {
		fmt.Println(n)
	}
}

func secIfPrac() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("n is 0")
	} else if n > 8 {
		fmt.Println("n is too big", n)
	} else {
		fmt.Println(n)
	}
}

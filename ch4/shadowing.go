package main

import "fmt"

func shadowingPrac() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func shadowingPrac2() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}

//func shadowingFmt() {
//	x := 10
//	fmt.Println(x)
//	fmt := "ooops"
//	fmt.Println(fmt)
//}

package main

import "fmt"

func SliceParc() {
	var x = []int{10, 20, 30}
	fmt.Println(x)
	var n []int
	fmt.Println(n)
	fmt.Println(n == nil)
	var appendedN = append(n, 10)
	fmt.Println(appendedN)
	var append2 = append(appendedN, 1, 2, 3)
	fmt.Println(append2)
	var spreadAppend = append(x, append2...)
	fmt.Println(spreadAppend)
}

func badSlice() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
}

func goodSlice() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
}

func copyPrac() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	y[1] = 20
	fmt.Println("y", y)
	fmt.Println("x", x)
	fmt.Println("num", num)
	z := make([]int, 3)
	copy(z[:2], x[1:])
	fmt.Println(z)
	num2 := copy(x[:3], x[1:])
	fmt.Println(x, num2)
}

func runeByteSlice() {
	var str = "hello 월드"
	fmt.Println(len(str))
	var rs []rune = []rune(str)
	var bs []byte = []byte(str)
	fmt.Println(rs)
	fmt.Println(bs)
}

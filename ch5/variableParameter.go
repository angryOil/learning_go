package main

import "fmt"

func variableParameterPrac() {
	fmt.Println(addTo(3, 4, 1, 2, 3, 4))
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(addTo(1, nums...))
}

func addTo(base int, values ...int) []int {
	out := make([]int, 0, len(values))
	for _, v := range values {
		out = append(out, base+v)
	}
	out = append(out, 0)
	fmt.Println(len(out), cap(out))
	return out
}

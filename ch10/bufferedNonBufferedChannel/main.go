package main

import "fmt"

func main() {
	var res = processChannel()
	fmt.Println(res)
}

func processChannel() []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			results <- i
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

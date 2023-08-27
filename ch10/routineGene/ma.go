package main

import "fmt"

func main() {
	//for i := range countTo(10) {
	//	fmt.Println(i)
	//}
	infinityAwait()
}

func infinityAwait() {
	for i := range countTo(10) {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

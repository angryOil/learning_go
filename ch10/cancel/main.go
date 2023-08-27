package main

import "fmt"

func main() {
	ch, cancel := countTo(10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()
}

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		fmt.Println("close")
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				fmt.Println("done")
				return
			default:
				fmt.Println("def")
				ch <- i
			}
		}
		// 반드시 반복문 밖에서 닫기
		close(ch)
	}()
	return ch, cancel
}

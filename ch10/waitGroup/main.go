package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println(doThing1())
	}()
	go func() {
		defer wg.Done()
		fmt.Println(doThing2())
	}()
	go func() {
		defer wg.Done()
		fmt.Println(doThing3())
	}()
	wg.Wait()
}

func doThing1() int {
	time.Sleep(1)
	return 1
}
func doThing2() int {
	time.Sleep(2)
	return 2
}
func doThing3() int {
	time.Sleep(3)
	return 3
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(initStart(1))
	fmt.Println(initStart(2))
	fmt.Println(initStart(3))
	fmt.Println(initStart(4))
	fmt.Println(initStart(5))
}

var once sync.Once

func initStart(num int) int {
	once.Do(func() {
		fmt.Println("init!!")
	})
	return num
}

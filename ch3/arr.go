package main

import "fmt"

func ArrPractice() {
	var x = [3]int{10, 20, 30}
	fmt.Println(x)
	var zeroInitArr [3]int
	fmt.Println(zeroInitArr)
	var someArr = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(someArr)
	var spreadArr = [...]int{1, 2, 3}
	fmt.Println(spreadArr)
	fmt.Println([3]int{1, 2, 3} == spreadArr)
}

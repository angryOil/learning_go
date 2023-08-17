package main

import "fmt"

func funcReturnFunc() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 4; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

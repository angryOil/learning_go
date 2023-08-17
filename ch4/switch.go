package main

import "fmt"

func switchPrac() {
	fir()
	exitSwitchLoopByLabel()
	blankSwitchTest()
}

func fir() {
	words := []string{"a", "cow", "world", "smile", "gopher", "moo ya ho", "warmOilW"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, " is a short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, " is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
			fallthrough
		default:
			fmt.Println(word, " is a long word")
		}
	}
}

func exitSwitchLoopByLabel() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "hello")
		case i%3 == 0:
			fmt.Println(i, "word")
		case i%7 == 0:
			fmt.Println(i, "exit loop")
			break loop
		}
	}
}

func blankSwitchTest() {
	words := []string{"hi", "hello", "world", "nice", "warmOil"}

	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen > 5:
			fmt.Println(word, " is to long")
		case wordLen < 5 && wordLen > 3:
			fmt.Println(word, " is to short")
		default:
			fmt.Println(word, " i don`t know ")
		}
	}
}

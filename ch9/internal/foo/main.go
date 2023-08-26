package main

import (
	"fmt"
	"learning_go/ch9/internal/foo/noInter"
)

func main() {
	noInter := noInter.NoInter{
		Name: "ads",
	}
	fmt.Println(noInter)
	//inter := Inter

}

package main

import (
	"fmt"
	"learning_go/ch9/internal/foo/internal"
)

func main() {
	inter := internal.Inter{
		Name: "helo",
	}
	fmt.Println(inter)
}

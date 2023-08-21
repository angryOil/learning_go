package main

import (
	"fmt"
	"sort"
)

func main() {
	dogs := Dogs{
		Dog{Name: "coco", Age: 11},
		Dog{Name: "kong", Age: 4},
		Dog{Name: "kongR", Age: 3},
	}
	sort.Sort(dogs)
	fmt.Println(dogs)
}

package main

import (
	"fmt"
	"sort"
)

func sortPrac() {
	sortTest()
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func sortTest() {
	people := []person{
		{"joy", "world", 22},
		{"angry", "oil", 21},
		{"warm", "oil", 23},
		{"jung", "onYu", 22},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println(people)
}

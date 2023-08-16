package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func practiceStruct() {
	var zeroPerson person
	literalZeroPerson := person{}
	fmt.Println(zeroPerson)
	fmt.Println(literalZeroPerson)

	var joy = person{
		name: "warmOil",
		age:  28,
		pet:  "coco",
	}
	fmt.Println(joy)
}

func useAnonymousStruct() {
	var anonymousStruct struct {
		name string
		age  int
		pet  string
	}
	anonymousStruct.name = "mooMung"
	anonymousStruct.age = 22
	anonymousStruct.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "coco",
		kind: "mixed",
	}

	fmt.Println(anonymousStruct)
	fmt.Println(pet)
}

func comparisonStruct() {
	type firstStruct struct {
		name string
		age  int
	}

	type secondStruct struct {
		name string
		age  int
	}

	var fir = firstStruct{
		name: "joy",
		age:  28,
	}
	var sec struct {
		name string
		age  int
	}
	sec = fir
	fmt.Println("::", sec, fir)
	fmt.Println(fir == sec)
}

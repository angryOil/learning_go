package main

import "fmt"

func deReferencing() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)
	//willBePanic()
	newKeyWord()
	forLiteralTypePointer()
	forPointerStruct()
	doFail()
	doUpdate()
}

// pointer를 역참조 하기 전에 nil 인지 확인하기
func willBePanic() {
	var x *int
	fmt.Println(x == nil)
	fmt.Println(*x)
}

func newKeyWord() {
	var x = new(int)
	fmt.Println(x == nil)
	print(*x)
}

func forLiteralTypePointer() {
	var y string
	z := &y
	fmt.Println(*z)
}

type pointerStruct struct {
	name       string
	strPointer *string
}

func getStringPointer(s string) *string {
	return &s
}

func forPointerStruct() {
	ps := pointerStruct{
		name:       "hello",
		strPointer: getStringPointer("hello"),
	}
	fmt.Println(ps)
	fmt.Println(*ps.strPointer)
}

func doUpdate() {
	x := 10
	failUpdate2(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}

func failUpdate2(p *int) {
	x2 := 20
	p = &x2
}

func update(p *int) {
	*p = 20
}

func doFail() {
	var f *int // f = nil
	fmt.Println(f)
	failedUpdate(f)
	fmt.Println(f)
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

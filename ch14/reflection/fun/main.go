package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
}

func main() {
	other()
}

func typeOf() {
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft)
	xpt := reflect.TypeOf(&xt)
	fmt.Println(">?:" + xpt.Name())
}

func other() {
	var x int
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())
	fmt.Println(xpt.Kind())
	fmt.Println(xpt.Elem().Name())
	fmt.Println(xpt.Elem().Kind())
}

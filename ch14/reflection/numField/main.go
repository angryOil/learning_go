package main

import (
	"fmt"
	"reflect"
)

func main() {
	useNumField()
}

type Foo struct {
	A int `myTag:"value"`
	B int `myTag:"value2"`
}

var f Foo

func useNumField() {
	ft := reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, ",", curField.Type.Name(), ",", curField.Tag.Get("myTag"))
	}
}

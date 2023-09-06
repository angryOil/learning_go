package main

import (
	"fmt"
	"reflect"
)

func main() {
	val()
	setByReflect()
	thatIsSameResult()
}

func val() {
	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	fmt.Println(sv)
	s2 := sv.Interface().([]string)
	fmt.Println(s2)
}

func setByReflect() {
	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(20)
	fmt.Println(i)
}

func thatIsSameResult() {
	i1 := 1
	i2 := 1
	chanInt(&i1)
	changeIntReflect(&i2)
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i1 == i2)
}

func chanInt(i *int) {
	*i = 20
}

func changeIntReflect(i *int) {
	iv := reflect.ValueOf(i)
	iv.Elem().SetInt(20)
}

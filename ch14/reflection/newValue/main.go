package main

import (
	"fmt"
	"reflect"
)

func main() {
	fir()
}

func fir() {
	var stringType = reflect.TypeOf((*string)(nil)).Elem()
	var stringSliceType = reflect.TypeOf([]string(nil))

	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	sv := reflect.New(stringType).Elem()
	sv.SetString("Hello")

	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss)
}

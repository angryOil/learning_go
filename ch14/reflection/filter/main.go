package main

import (
	"fmt"
	"reflect"
)

func main() {
	useFilter()
}

func useFilter() {
	names := []string{"andrew", "bob", "clara", "hortense"}
	longNames := filter(names, func(s string) bool {
		return len(s) > 3
	}).([]string)
	fmt.Println(longNames)

	ages := []int{20, 40, 30, 14, 12}
	adults := filter(ages, func(a int) bool {
		return a >= 20
	}).([]int)
	fmt.Println(adults)

}

func filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		}
	}
	return out.Interface()
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	useCustomThings()
}
func useCustomThings() {
	var mi myInt = 22
	fmt.Println(mi.String())
	var ami alsoMyInt = 333
	fmt.Println(ami.String())
	var mf myFloat = 44.2
	fmt.Print(mf.String())
}

type myInt int

func (mi myInt) String() string {
	return strconv.Itoa(int(mi))
}

type alsoMyInt int

func (ami alsoMyInt) String() string {
	return "42"
}

type myFloat float64

func (mf myFloat) String() string {
	return strconv.FormatFloat(float64(mf), 'E', -1, 64)
}

//func useRangeAble() {
//	//x := []string{"a", "b", "c"}
//	//var out []string = copyVals[interface{}, string, []string](x)
//	//fmt.Println(out)
//
//	y := map[string]float64{"a": 1.2, "b": 3.1, "c": 10}
//	var out2 []float64 = copyVals[string, float64, map[string]float64](y)
//	fmt.Print(out2)
//}

type rangeAble[E comparable, T any] interface {
	[]T | map[E]T
}

//func copyVals[E comparable, T any, R rangeAble[E, T]](vals R) []T {
//	var out []T
//	for _, v := range vals {
//		out = append(out, v)
//	}
//	return out
//}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mti myTypes
	mti = myInt(22)
	fmt.Print(mti)
}

type myTypes interface {
	String() string
}

type myInt int

func (mi myInt) String() string {
	return strconv.Itoa(int(mi))
}

type myFloat float64

func (mf myFloat) String() string {
	return strconv.FormatFloat(float64(mf), 'E', -1, 64)
}

package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	meTest()
}

func meTest() {
	err := AnfunctionThatReturnAnError()
	var me MyErr
	if errors.As(err, &me) {
		fmt.Println(me.Codes)
	}
}

func AnfunctionThatReturnAnError() error {
	return MyErr{Codes: []int{1, 2, 3, 45}}
}

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes:%v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if m2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, m2)
	}
	return false
}

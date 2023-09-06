package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	modifyStr()
	useSliceHeader()
}

func useSliceHeader() {
	s := []int{10, 20, 30}
	sHdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len)
	fmt.Println(sHdr.Cap)
	fmt.Println(sHdr.Data)

	intByteSize := unsafe.Sizeof(s[0])
	fmt.Println(intByteSize)
	for i := 0; i < sHdr.Len; i++ {
		intVal := *(*int)(unsafe.Pointer(sHdr.Data + intByteSize*uintptr(i)))
		fmt.Println(intVal)
	}
}

func modifyStr() {
	s := "original"
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len)
	for i := 0; i < sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(sHdr.Data + uintptr(i)))
		fmt.Println(bp)
	}
	fmt.Println()
	runtime.KeepAlive(s)
}

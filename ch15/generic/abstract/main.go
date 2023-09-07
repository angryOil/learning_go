package main

import "fmt"

func main() {
	useFunc()
	useConvert()
	fmt.Println(plusTenThousand[int64](int64(2)))
}

// 안되는코드
//func plusOneThousand[T integer](in integer) T {
//	return in + 100
//}

func plusTenThousand[T integer](in T) T {
	return in + 127
}

func useConvert() {
	var a = -1
	b := convert[int, uint](a)
	fmt.Println(b)
}

func useFunc() {
	words := []string{"one", "two", "three", "world", "hello"}
	filteredWord := Filter(words, func(s string) bool {
		return s != "three"
	})
	fmt.Println(filteredWord)

	lengths := Map(filteredWord, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)

	sum := Reduce(lengths, 0, func(acc, val int) int {
		return acc + val
	})
	fmt.Println(sum)
}

func Map[T, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T, R any](s []T, initializer R, f func(R, T) R) R {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

type integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func convert[T, R integer](in T) R {
	return R(in)
}

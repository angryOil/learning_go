package bench

import (
	"fmt"
	"os"
	"testing"
)

func TestFileLen(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dir)
	result, err := FileLen("test.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

var blackhole int

func BenchmarkFileLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := FileLen("test.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = result
	}
}

func BenchmarkFileLen2(b *testing.B) {
	for _, v := range []int{1, 10, 100, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("fileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("test.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}

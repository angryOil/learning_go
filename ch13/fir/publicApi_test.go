package fir_test

// packageName_test 형태로 패키지 지정시 공개 api만 호출가능
import (
	"learning_go/ch13/fir"
	"testing"
)

func TestMulNumbers(t *testing.T) {
	num := fir.MulNumbers(2, 3)
	if num != 6 {
		t.Error("mulNumbers(2,3) expected 6 but was:", num)
	}
}

package fir

import "testing"

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result : expected  5, got", result)
		t.Errorf("incorrect result: expected %d ,got %d", 5, result)
	}
}

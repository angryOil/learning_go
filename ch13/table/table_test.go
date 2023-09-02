package table

import "testing"

// 놀랍다!!
var data = []struct {
	name     string
	num1     int
	num2     int
	op       string
	expected int
	errMsg   string
}{
	{"addition", 2, 2, "+", 4, ""},
	{"subtraction", 2, 2, "-", 0, ""},
	{"multiplication", 2, 2, "*", 4, ""},
	{"division", 2, 2, "/", 1, ""},
	{"div by zero", 2, 0, "/", 0, `division by zero`},
	{"bad_op", 2, 2, "?", 0, "unknown operator ?"},
}

func TestDoMath(t *testing.T) {
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expect %d , god %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("expected error message %s , got %s", d.errMsg, errMsg)
			}
		})
	}
}

package stup

import (
	"context"
)

type Processor struct {
	Solver MathSolver
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
}

//func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
//	curExpression, err := readToNewLine()
//}

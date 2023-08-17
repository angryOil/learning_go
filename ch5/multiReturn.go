package main

import (
	"errors"
	"fmt"
	"os"
)

func multiReturnPrac() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
	fmt.Println(blankReturn(2, 3))
}

func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func namedReturnDivAndRemainder(numerator, denominator int) (result, remainder int, err error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func blankReturn(num, den int) (result, remainder int, err error) {
	if den == 0 {
		err = errors.New("den cannot be 0 ")
		return
	}
	result, remainder = num/den, num%den
	return

}

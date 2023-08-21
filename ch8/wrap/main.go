package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	unwrapPrac()
}

func unwrapPrac() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

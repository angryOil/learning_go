package main

import (
	"errors"
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	err     error
}

func vUserInsteadW() error {
	err := StatusErr{
		Status:  0,
		Message: "msg",
		err:     errors.New("hello"),
	}
	if err.Unwrap() != nil {
		return fmt.Errorf("internal filaure: %v", err.Unwrap().Error())
	}
	return nil
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.err
}

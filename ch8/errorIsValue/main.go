package main

import "fmt"

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
	err = GenerateError2(true)
	fmt.Println(err != nil)
	err = GenerateError2(false)
	fmt.Println(err != nil)
	err = GenerateError3(true)
	fmt.Println(err != nil)
	err = GenerateError3(false)
	fmt.Println(err != nil)

}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func GenerateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func GenerateError2(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

func GenerateError3(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

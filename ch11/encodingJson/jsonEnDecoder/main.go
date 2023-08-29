package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	encFileToTemp()
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func encFileToTemp() {
	toFile := Person{
		Name: "red",
		Age:  40,
	}

	tempFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())
	err = json.NewEncoder(tempFile).Encode(toFile)
	if err != nil {
		panic(err)
	}
	err = tempFile.Close()
	if err != nil {
		panic(err)
	}

	tempFile2, err := os.Open(tempFile.Name())
	if err != nil {
		panic(err)
	}
	var fromFile Person
	err = json.NewDecoder(tempFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}
	err = tempFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(fromFile)
}

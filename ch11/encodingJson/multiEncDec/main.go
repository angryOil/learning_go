package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	decode()
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var data = `
	{"name":"hello","age":22}
	{"name":"hello2","age":12}
	{"name":"hell3","age":23}
	{"name":"he4lo","age":25}
`

func decode() {
	dec := json.NewDecoder(strings.NewReader(data))
	var t Person
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)
	}
}

func encode() {
	var b bytes.Buffer
	inputs := []Person{{
		Name: "22",
		Age:  1,
	}, {
		Name: "asd",
		Age:  231,
	}}
	enc := json.NewEncoder(&b)
	for _, input := range inputs {
		enc.Encode(input)
	}
}

package main

import "fmt"

func duckTyping() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}

type LogicProvider struct {
}

func (lp LogicProvider) Process(data string) string {
	return data
}

type Logic2 interface {
	Process(data string) string
}

type Client struct {
	L Logic2
}

func (c Client) Program() {
	fmt.Println(c.L.Process("hello"))
}

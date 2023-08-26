package main

import "fmt"

func main() {
	var bar = makeBar()
	fmt.Println(bar.goodBye())
	bar.newFunc()
	var foo = Foo{
		x: 21,
		S: "hello?",
	}
	foo.newFunc()
}

type Foo struct {
	x int
	S string
}

func (f Foo) hello() string {
	return "hello world"
}

func (f Foo) goodBye() string {
	return "good bye"
}

type Bar = Foo

func (b Bar) newFunc() {
	fmt.Println(b.goodBye() + "new")
}

func makeBar() Bar {
	bar := Bar{
		x: 20,
		S: "hello",
	}
	var f Foo = bar
	fmt.Println(f.hello())
	fmt.Println(bar.goodBye())
	return f
}

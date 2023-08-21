package main

type Dog struct {
	Name string
	Age  int
}

type Dogs []Dog

func (d Dogs) Len() int           { return len(d) }
func (d Dogs) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d Dogs) Less(i, j int) bool { return len(d[i].Name) < len(d[j].Name) }

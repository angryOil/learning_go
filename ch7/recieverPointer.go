package main

import (
	"fmt"
	"time"
)

func useRecieverPointer() {
	var counter = Counter{
		total: 1,
	}
	fmt.Println(counter.String())
	counter.Increment()
	fmt.Println(counter.String())

	var counter2 Counter
	doUpdateWrong(counter2)
	fmt.Println(counter2)
	doUpdateRight(&counter2)
	fmt.Println(counter2)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d , last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

package main

import "time"

func main() {
	err := parse()
	if err != nil {
		print(err)
	}
}

func parse() error {
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		return err
	}
	print(t.Format("January 2, 2006 at 3:305:05PM MST"))
	return nil
}

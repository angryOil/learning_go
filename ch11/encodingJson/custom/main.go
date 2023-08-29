package main

import (
	"fmt"
	"time"
)

func main() {
	o := Order{
		id:       "asd",
		datetime: RFC822ZTime{time.Now()},
	}
	fmt.Println(o)
	fmt.Println(o.datetime)
}

type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC822ZTime{t}
	return nil
}

type Order struct {
	id       string      `json:"id"`
	datetime RFC822ZTime `json:"datetime"`
}

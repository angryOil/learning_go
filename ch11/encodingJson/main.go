package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	unmarshal()
	marshal()
}

type Oder struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Items      []Item `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func unmarshal() {
	mockJson := `{
	"id" :"dds",
	"customer_id":"3431a",
	"items":[{"id":"id1","name":"name1"},{"id":"id2","name":"name2"}]
}`
	var o Oder
	err := json.Unmarshal([]byte(mockJson), &o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o)
}

func marshal() {
	o := Oder{
		ID:         "123",
		CustomerID: "aa1",
		Items: []Item{{
			ID:   "a",
			Name: "b",
		}, {
			ID:   "c",
			Name: "d",
		},
		},
	}
	json, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	createRequest()
}

var data struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func createRequest() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	fmt.Println(client)
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, "https://jsonplaceholder.typicode.com/todos/12", nil,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-MY-Client", "Leaning GO")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}

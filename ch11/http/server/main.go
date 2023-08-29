package main

import (
	"net/http"
	"time"
)

func main() {
	registerMuxHandler()
	registerHelloHandler()
}

func registerHelloHandler() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      HelloHandler{},
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func registerMuxHandler() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mux", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("well come to mux handle"))
	})
	http.ListenAndServe(":8080", mux)
}

type HelloHandler struct {
}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("fine??"))
}

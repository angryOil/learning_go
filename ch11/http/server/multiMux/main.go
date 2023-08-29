package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	registerHandler()
}

func registerHandler() {
	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello this is person handler"))
	})
	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is dog handler"))
	})
	mux := http.NewServeMux()
	timer := requestTimer
	mux.Handle("/person/", timer(http.StripPrefix("/person", person)))
	mux.Handle("/dog/", timer(http.StripPrefix("/dog", dog)))

	http.ListenAndServe(":8080", mux)
}

func requestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("reuqest time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

package main

import "fmt"

func main() {
	searchers := []func(string) []string{
		searchGoogle,
		searchBing,
		searchDuckDuckGo,
	}

	searchResults := searchData("programming", searchers)
	fmt.Println("Search Results:", searchResults)
}

func searchGoogle(query string) []string {
	fmt.Println("Searching Google for:", query)
	return []string{"Google Result 1", "Google Result 2"}
}

func searchBing(query string) []string {
	fmt.Println("Searching Bing for:", query)
	return []string{"Bing Result 1", "Bing Result 2"}
}

func searchDuckDuckGo(query string) []string {
	fmt.Println("Searching DuckDuckGo for:", query)
	return []string{"DuckDuckGo Result 1", "DuckDuckGo Result 2"}
}

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:

			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}

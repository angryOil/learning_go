package main

import "context"

func main() {
	ss := slowServer()
	defer ss.Close()
	fs := fastServer()
	defer fs.Close()

	ctx := context.Background()
	callBoth(ctx, "false", ss.URL, fs.URL)
	//callBoth(ctx, "true", ss.URL, fs.URL)
}

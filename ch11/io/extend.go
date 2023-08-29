package main

import "io"

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

func NopClose(r io.Reader) io.ReadCloser {
	return nopCloser{r}
}

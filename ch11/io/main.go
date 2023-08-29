package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	useStringIo()
	//makeGZip()
	useGZip()
}

var baseGZipName = "ch11/io/hello.txt.gz"

func makeGZip() {
	sourceFilePath := "ch11/io/hello.txt"
	compressedFilePath := "ch11/io/hello.txt.gz"
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer sourceFile.Close()

	compressedFile, err := os.Create(compressedFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer compressedFile.Close()

	gzipWriter := gzip.NewWriter(compressedFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, sourceFile)
	if err != nil {
		fmt.Println(err)
	}
}

func useStringIo() {
	s := "the quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(counts)
}

func useGZip() {
	r, closer, err := buildGZipReader()
	if err != nil {
		panic(err)
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(counts)
}

func buildGZipReader() (*gzip.Reader, func(), error) {
	r, err := os.Open(baseGZipName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func useCountLetters() {
	content, err := os.ReadFile("ch11/io/hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	var fileStr = bytes.NewReader(content)
	result, err := countLetters(fileStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err != io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

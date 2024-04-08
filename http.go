package main

import (
	"fmt"
	"io"
	"net/http"
)

type Writer struct{}

func main() {
	fmt.Println("hello")
	wr := Writer{}
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(wr, resp.Body)
}

func (Writer) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return 0, nil
}

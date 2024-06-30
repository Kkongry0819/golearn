package main

import (
	"fmt"
	"io"
	"net/http"
)

func retrieve(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body), nil
}
func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var s string
	s, _ = retrieve("http://www.baidu.com")
	fmt.Printf("%s", s)
}

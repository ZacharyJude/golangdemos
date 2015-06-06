package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	url, _ := url.Parse("http://www.baidu.com")
	resp, _ := http.Get(url.String())
	for k, v := range resp.Header {
		fmt.Printf("%s\t%s\n", k, v)
	}

	fmt.Printf("====%s\n", resp.Header.Get("Server"))
	fmt.Printf("====%s\n", resp.Header.Get("server"))
}

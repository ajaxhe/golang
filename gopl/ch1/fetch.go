package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//fetch(resp, url)
		fetch2(resp, url)
	}
}

func fetch(resp *http.Response, url string) int {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Printf("%s", b)
	return 0
}

func fetch2(resp *http.Response, url string) int {
	bytes, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Printf("url:%s\n", url)
	fmt.Printf("resp status:%s, resp code: %d, read bytes:%d\n", resp.Status, resp.StatusCode, bytes)
	return 0
}

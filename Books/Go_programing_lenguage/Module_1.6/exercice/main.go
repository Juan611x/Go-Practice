package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	f, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for range os.Args[1:] {
		fmt.Fprint(f, <-ch) // receive from channel
	}
	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v\n", err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetch: reading response body: %v\n", err)
		return
	}
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", time.Since(start).Seconds(), nbytes, url)
}

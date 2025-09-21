package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// go run .\Books\Go_programing_lenguage\Module_1.5\fetch.go http://gopl.io

func formatUrl(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	return url
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(formatUrl(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Content: \n")
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading response body: %v\n", err)
			os.Exit(1)
		}
	}
}

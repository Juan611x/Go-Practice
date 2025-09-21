package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// go run .\Books\Go_programing_lenguage\Module_1.5\fetch.go http://gopl.io

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading response body: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

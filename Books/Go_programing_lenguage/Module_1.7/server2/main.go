package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// go run .\Books\Go_programing_lenguage\Module_1.7\server2\main.go
// go run .\Books\Go_programing_lenguage\Module_1.5\fetch\main.go http://localhost:8000/hola
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()	
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

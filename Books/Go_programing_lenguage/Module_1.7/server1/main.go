package main

import (
	"fmt"
	"net/http"
)

// go run .\Books\Go_programing_lenguage\Module_1.7\server1\main.go
// go run .\Books\Go_programing_lenguage\Module_1.5\fetch\main.go http://localhost:8000/hola
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

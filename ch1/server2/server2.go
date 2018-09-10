package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", hander)
	http.HandleFunc("/count", counter)
	http.ListenAndServe("localhost:8000", nil)
}

func hander(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "count=%d URL.Path=%q \n", count, r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count=%d\n", count)
	mu.Unlock()
}

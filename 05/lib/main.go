package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

type hey struct {
	n int
}

func (h2 *hey) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func main() {
	http.Handle("/count", new(countHandler))
	http.Handle("/count2", new(hey))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

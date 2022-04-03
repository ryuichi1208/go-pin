package main

import (
	"fmt"
	"net/http"
)

type Inter interface {
	Next(string) error
}

type Istr struct {
}

func (i Istr) Next(s string) error {
	return nil
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}

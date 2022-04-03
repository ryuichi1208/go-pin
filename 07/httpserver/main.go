package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	fmt.Fprintf(w, strconv.Itoa(rand.Intn(100)))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

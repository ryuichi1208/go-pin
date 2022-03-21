package main

import (
	"fmt"
	"net/http"
)

type Human struct {
	age  int
	name string
	sex  *string
}

func (h *Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", *h.sex)

}

func main() {
	fmt.Println("hello")
	sex := "m"
	h := &Human{age: 10, name: "ri", sex: &sex}
	http.Handle("/", h)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

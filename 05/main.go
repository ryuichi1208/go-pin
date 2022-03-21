package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func hand2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	fmt.Println(r.Header)
}

type Page struct {
	Title string
	Count int
}

type H struct {
	a string
}

func (h *H) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aaa")
}

func viweHandler(w http.ResponseWriter, r *http.Request) {
	page := &Page{Title: "test", Count: 3}
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type hey string

func (h hey) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "aaa")
}

func main() {
	fmt.Println(String("aaa"))
	http.Handle("/", String("Hello World."))
	http.Handle("/s", new(hey))
	http.ListenAndServe("localhost:8000", nil)
}

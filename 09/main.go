package main

import "fmt"

type Hi interface {
	h(string) string
}

type Hs struct{}

func (h Hs) h(a string) string {
	return "prod" + a
}

func run(hi Hi) string {
	return hi.h("aaa")
}

func main() {
	t := &Hs{}
	fmt.Println(run(t))
}

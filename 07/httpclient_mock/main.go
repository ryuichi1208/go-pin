package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIer interface {
	DoHttpRequest(url string) string
	DoHttpRequest2(url string) string
	DoHttpRequest3(url string) string
	DoHttpRequest4(url string) string
	DoHttpRequest5(url string) string
}

type API struct {
}

func (a API) DoHttpRequest(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}
func (a API) DoHttpRequest2(url string) string {
	return string("1")
}
func (a API) DoHttpRequest3(url string) string {
	return string("1")
}
func (a API) DoHttpRequest4(url string) string {
	return string("1")
}
func (a API) DoHttpRequest5(url string) string {
	return string("1")
}

func Run(s string, a APIer) string {
	return s + a.DoHttpRequest("http://localhost:8080")
}

func main() {
	a := &API{}
	fmt.Println(Run("test: ", a))
}

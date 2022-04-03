package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoHttpRequest(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(byteArray)) // htmlをstringで取得

	return string(byteArray)
}

func Run(s string) string {
	return s + DoHttpRequest("http://localhost:8080")
}

func main() {
	fmt.Println(Run("test: "))
}

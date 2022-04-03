package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mock hello"))
}

func TestRoot(t *testing.T) {
	// testserver := httptest.NewServer(http.HandlerFunc(mockHelloHandler))
	testserver := httptest.NewServer(http.HandlerFunc(HelloHandler))
	defer testserver.Close()
	fmt.Println(testserver.URL)
	res, err := http.Get(testserver.URL)
	if err != nil {
		t.Error(err)
	}
	hello, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("not 200")
	}
	if string(hello) != "hello" {
		t.Errorf("not string")
	}

}

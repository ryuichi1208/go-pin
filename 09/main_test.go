package main

import "testing"

type Hmock struct{}

func (h Hmock) h(a string) string {
	return "mock" + a
}

func TestHmock(t *testing.T) {
	hm := &Hmock{}
	s := run(hm)

	if s != "mockaaa" {
		t.Error(s)
	}
}

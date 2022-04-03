package main

import (
	"testing"
)

type APIMock struct {
	APIer
}

func (a APIMock) DoHttpRequest(url string) string {
	return ""
}

func TestRand(t *testing.T) {
	ai := &APIMock{}
	tests := []struct {
		a string
		b string
	}{
		{
			a: "test: 1",
			b: "test: 1",
		},
	}

	for _, tt := range tests {
		res := Run(tt.a, ai)
		if tt.b != res {
			t.Errorf("got: %s, want: %s", res, tt.b)
		}
	}
}

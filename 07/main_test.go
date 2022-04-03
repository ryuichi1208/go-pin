package main

import "testing"

type RandomNumMock struct {
	RandomNumer
	ReturnRandomNumMock func(int) int
}

func (r RandomNumMock) ReturnRandomNum(n int) int {
	return r.ReturnRandomNumMock(n)
}

func TestMock(t *testing.T) {
	rnm := &RandomNumMock{
		ReturnRandomNumMock: func(n int) int {
			return n
		},
	}

	res := rnm.ReturnRandomNum(1)
	if res != 1 {
		t.Errorf("want 2, get %d\n", res)
	}
}

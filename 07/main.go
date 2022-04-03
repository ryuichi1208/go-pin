package main

import (
	"math/rand"
	"time"
)

type RandomNumer interface {
	ReturnRandomNum(int) int
}

type RandomNum struct {
	r RandomNumer
}

func (r RandomNum) ReturnRandomNum(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

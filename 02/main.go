package main

import (
	"log"
	"math"
	"sync"
	"time"
)

const concurrency = 4

func add(k, v string) map[string]int {
	a := map[string]int{"aaa": 10}
	return a
}

func main() {
	var wg sync.WaitGroup
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	sem := make(chan struct{}, concurrency)
	for _, n := range nums {
		sem <- struct{}{}
		wg.Add(1)
		go func(n int) {
			time.Sleep(2 * time.Second)
			log.Println(n, math.Pi)
			wg.Done()
			<-sem
		}(n)
	}

	wg.Wait()
}

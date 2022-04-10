package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f(msg string, n int) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < n; i++ {
			ch <- msg + " please!"
		}
		wg.Done()
	}()
	return ch
}

func main() {
	wg.Add(3)
	ch1 := f("beer", 4)
	ch2 := f("juice", 3)
	ch3 := f("wine", 2)

	done := make(chan bool)

	go func() {
		wg.Wait()
		done <- true
	}()

L:
	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case msg := <-ch3:
			fmt.Println(msg)
		case <-done:
			break L
		}
	}
	fmt.Println("exit")

}

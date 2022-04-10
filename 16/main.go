package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	filename := "./text.txt"
	rand.Seed(time.Now().UnixNano())
	for range time.Tick(3 * time.Second) {
		text := fmt.Sprintf("%d", rand.Intn(100))
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)

		}

		defer f.Close()

		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
		log.Println(text)
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/04/lib/linux"
)

func main() {
	fmt.Println(1)
	time.Sleep(3 * time.Second)
	linux.Do()
}

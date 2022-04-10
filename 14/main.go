package main

import "fmt"

func main() {
	i := 0
	switch {
	case i == 0, i == 1:
		fmt.Println(0)
		fallthrough
	case i == 2, i == 3:
		fmt.Println(2)
	default:
		fmt.Println(1)
	}
}

package main

import (
	"fmt"
	"net"
)

type A struct {
	a *string
}

type B struct {
	d *A
}

type C struct {
	B
}

type D interface{}

func main() {
	if err := run(); err != nil {
		fmt.Println("%+v", err)
	}
}

func run() error {
	fmt.Println("start tcp listen...")
	listen, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println(">>> start")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
	}
	fmt.Println("<<< end")
	return nil
}

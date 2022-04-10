package cmd

import (
	"fmt"
	"net"
)

func Do() int {
	addr, err := net.ResolveIPAddr("ip", "www.google.com")
	if err != nil {
		return 1
	}

	fmt.Println("Resolve", addr.String())
	return 0
}

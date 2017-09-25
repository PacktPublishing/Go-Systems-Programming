package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an IP address!")
		os.Exit(100)
	}

	IP := arguments[1]
	addr := net.ParseIP(IP)
	if addr == nil {
		fmt.Println("Not a valid IP address!")
		os.Exit(100)
	}

	hosts, err := net.LookupAddr(IP)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	for _, hostname := range hosts {
		fmt.Println(hostname)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(100)
	}

	hostname := arguments[1]
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	for _, IP := range IPs {
		fmt.Println(IP)
	}
}

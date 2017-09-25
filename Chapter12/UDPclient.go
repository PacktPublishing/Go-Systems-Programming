package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		os.Exit(100)
	}
	CONNECT := arguments[1]

	s, err := net.ResolveUDPAddr("udp", CONNECT)
	c, err := net.DialUDP("udp", nil, s)

	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	data := []byte("Hello UDP Echo server!\n")
	_, err = c.Write(data)

	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	fmt.Print("Reply: ", string(buffer[:n]))
}

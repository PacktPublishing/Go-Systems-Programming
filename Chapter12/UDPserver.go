package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		os.Exit(100)
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	connection, err := net.ListenUDP("udp", s)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n]))
		data := []byte(buffer[0:n])
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}
	}
}

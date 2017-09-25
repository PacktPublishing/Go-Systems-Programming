package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a server:port string!")
		os.Exit(100)
	}

	CONNECT := arguments[1]
	myMessage := "Hello from TCP client!\n"

	tcpAddr, err := net.ResolveTCPAddr("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	_, err = conn.Write([]byte(myMessage))
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	fmt.Print("-> ", myMessage)
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	fmt.Print(">> ", string(buffer[0:n]))
	conn.Close()
}

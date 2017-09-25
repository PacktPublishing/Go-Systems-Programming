package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		fmt.Print("-> ", string(netData))
		c.Write([]byte(netData))
		if strings.TrimSpace(string(netData)) == "STOP" {
			break
		}
	}
	time.Sleep(3 * time.Second)
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		os.Exit(100)
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		go handleConnection(c)
	}
}

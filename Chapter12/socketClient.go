package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func readSocket(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("-> ", string(buf[0:n]))
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a socket file.")
		os.Exit(100)
	}
	socketFile := arguments[1]

	c, err := net.Dial("unix", socketFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	defer c.Close()

	go readSocket(c)
	for {
		_, err := c.Write([]byte("Hello Server!"))
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		time.Sleep(1 * time.Second)
	}
}

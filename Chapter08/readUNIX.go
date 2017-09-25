package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func readSocket(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, _ := r.Read(buf[:])
		fmt.Print("Read: ", string(buf[0:n]))
	}
}

func main() {
	c, _ := net.Dial("unix", "/tmp/aSocket.sock")
	defer c.Close()

	go readSocket(c)
	n := 0
	for {
		message := []byte("Hi there: " + strconv.Itoa(n) + "\n")
		_, _ = c.Write(message)
		time.Sleep(5 * time.Second)
		n = n + 1
	}
}

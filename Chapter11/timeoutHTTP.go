package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var timeout = time.Duration(time.Second)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	URL := os.Args[1]
	t := http.Transport{
		Dial: Timeout,
	}

	client := http.Client{
		Transport: &t,
	}
	data, err := client.Get(URL)

	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
	}
}

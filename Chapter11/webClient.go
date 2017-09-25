package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Parse:", err)
		os.Exit(100)
	}

	c := &http.Client{}

	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	fmt.Println("Status code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	fmt.Println("Character Set:", characterSet[1])

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength in unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Response data length:", length)
}

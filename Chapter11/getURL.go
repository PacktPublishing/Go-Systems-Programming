package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	URL := os.Args[1]
	data, err := http.Get(URL)

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

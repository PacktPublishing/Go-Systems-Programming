package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s [-t|-s] filename!\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	convertTabs := false
	convertSpaces := false
	newLine := ""

	option := os.Args[1]
	filename := os.Args[2]
	if option == "-t" {
		convertTabs = true
	} else if option == "-s" {
		convertSpaces = true
	} else {
		fmt.Println("Unknown option!")
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening %s: %s", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			os.Exit(1)
		}

		if convertTabs == true {
			newLine = strings.Replace(line, "\t", "    ", -1)
		} else if convertSpaces == true {
			newLine = strings.Replace(line, "    ", "\t", -1)
		}
		fmt.Print(newLine)
	}
}

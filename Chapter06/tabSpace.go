package main

import (
	"os"
	"fmt"
	"path/filepath"
	"bufio"
	"io"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s [-t|-s] filename!\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	option := os.Args[1]
	filename := os.Args[2]
	if option != "-t" && option != "-s" {
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
			outputNewLine(option, line)
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			os.Exit(1)
		}

		outputNewLine(option,line)
	}
}

func outputNewLine(option, line string) {
	if option == "-t" {
		fmt.Print(strings.Replace(line, "\t", "    ", -1))
	}else {
		fmt.Print(strings.Replace(line, "    ", "\t", -1))
	}
}
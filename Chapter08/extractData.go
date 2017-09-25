package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Printf("%s IP <files>\n", filepath.Base(os.Args[0]))
		os.Exit(-1)
	}

	WANTED := arguments[1]
	for _, filename := range arguments[2:] {
		count := 0
		buf := []byte(filename)
		io.WriteString(os.Stdout, string(buf))
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error in file: %s\n", err)
				continue
			}

			ip := findIP(line)
			if ip == WANTED {
				count = count + 1
			}
		}
		buf = []byte(strconv.Itoa(count))
		io.WriteString(os.Stdout, " ")
		io.WriteString(os.Stdout, string(buf))
		io.WriteString(os.Stdout, "\n")
	}
}

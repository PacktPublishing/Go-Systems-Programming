package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	minusCOL := flag.Int("COL", 1, "Column")
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Printf("usage: readColumn <file1> [<file2> [... <fileN]]\n")
		os.Exit(1)
	}

	column := *minusCOL

	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	for _, filename := range flags {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}

			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println((data[column-1]))
			}
		}
	}
}

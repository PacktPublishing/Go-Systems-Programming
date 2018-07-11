package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	minusINIT := flag.Int("init", 1, "Initial Value")
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Printf("usage: insertLineNumber <file1> [<file2> [... <fileN]]\n")
		os.Exit(1)
	}

	lineNumber := *minusINIT
	for _, filename := range flags {
		fmt.Println("Processing:", filename)

		input, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		lines := strings.Split(string(input), "\n")
		for i, line := range lines {
			lines[i] = fmt.Sprintf("%d: %s ", lineNumber, line)
			lineNumber = lineNumber + 1
		}

		lines[len(lines)-1] = ""
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	fmt.Println("Processed", lineNumber-*minusINIT, "lines!")
}

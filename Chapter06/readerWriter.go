package main

import (
	"fmt"
	"io"
	"os"
)

func countChars(r io.Reader) int {
	buf := make([]byte, 16)
	total := 0
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0
		}
		if err == io.EOF {
			break
		}
		total = total + n
	}
	return total
}

func writeNumberOfChars(w io.Writer, x int) {
	fmt.Fprintf(w, "%d\n", x)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Error opening %s.\n", filename)
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file:", err)
		os.Exit(-1)
	}
	defer f.Close()

	chars := countChars(f)
	filename = filename + ".Count"
	f, err = os.Create(filename)
	if err != nil {
		fmt.Println("os.Create:", err)
		os.Exit(1)
	}
	defer f.Close()
	writeNumberOfChars(f, chars)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type File struct {
	Filename   string
	Lines      int
	Words      int
	Characters int
	Error      error
}

func monitor(values <-chan File, count int) {
	var totalWords int = 0
	var totalLines int = 0
	var totalChars int = 0
	for i := 0; i < count; i++ {
		x := <-values
		totalWords = totalWords + x.Words
		totalLines = totalLines + x.Lines
		totalChars = totalChars + x.Characters
		if x.Error == nil {
			fmt.Printf("\t%d\t", x.Lines)
			fmt.Printf("%d\t", x.Words)
			fmt.Printf("%d\t", x.Characters)
			fmt.Printf("%s\n", x.Filename)
		} else {
			fmt.Printf("\t%s\n", x.Error)
		}
	}
	fmt.Printf("\t%d\t", totalLines)
	fmt.Printf("%d\t", totalWords)
	fmt.Printf("%d\ttotal\n", totalChars)
}

func count(filename string, out chan<- File) {
	var err error
	var nLines int = 0
	var nChars int = 0
	var nWords int = 0

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		newValue := File{Filename: filename, Lines: 0, Characters: 0, Words: 0, Error: err}
		out <- newValue
		return
	}

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s\n", err)
		}
		nLines++
		r := regexp.MustCompile("[^\\s]+")
		for range r.FindAllString(line, -1) {
			nWords++
		}
		nChars += len(line)
	}
	newValue := File{Filename: filename, Lines: nLines, Characters: nChars, Words: nWords, Error: nil}
	out <- newValue
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	values := make(chan File, len(os.Args[1:]))
	for _, filename := range os.Args[1:] {
		go func(filename string) {
			count(filename, values)
		}(filename)
	}
	monitor(values, len(os.Args[1:]))
}

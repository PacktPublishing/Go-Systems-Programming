package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type File struct {
	Filename   string
	Lines      int
	Words      int
	Characters int
	Error      error
}

var aM sync.Mutex
var values = make([]File, 0)

func count(filename string) {
	var err error
	var nLines int = 0
	var nChars int = 0
	var nWords int = 0

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		newValue := File{Filename: filename, Lines: 0, Characters: 0, Words: 0, Error: err}
		aM.Lock()
		values = append(values, newValue)
		aM.Unlock()
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
	aM.Lock()
	values = append(values, newValue)
	aM.Unlock()
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	var waitGroup sync.WaitGroup
	for _, filename := range os.Args[1:] {
		waitGroup.Add(1)
		go func(filename string) {
			count(filename)
			defer waitGroup.Done()
		}(filename)
	}

	waitGroup.Wait()

	var totalWords int = 0
	var totalLines int = 0
	var totalChars int = 0
	for _, x := range values {
		totalWords = totalWords + x.Words
		totalLines = totalLines + x.Lines
		totalChars = totalChars + x.Characters
		if x.Error == nil {
			fmt.Printf("\t%d\t", x.Lines)
			fmt.Printf("%d\t", x.Words)
			fmt.Printf("%d\t", x.Characters)
			fmt.Printf("%s\n", x.Filename)
		}
	}
	fmt.Printf("\t%d\t", totalLines)
	fmt.Printf("%d\t", totalWords)
	fmt.Printf("%d\ttotal\n", totalChars)
}

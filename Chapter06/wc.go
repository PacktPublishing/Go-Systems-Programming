package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func countLines(filename string) (int, int, int) {
	var err error
	var numberOfLines int
	var numberOfCharacters int
	var numberOfWords int
	numberOfLines = 0
	numberOfCharacters = 0
	numberOfWords = 0

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", err)
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
		}

		numberOfLines++
		r := regexp.MustCompile("[^\\s]+")
		for range r.FindAllString(line, -1) {
			numberOfWords++
		}
		numberOfCharacters += len(line)
	}

	return numberOfLines, numberOfWords, numberOfCharacters
}

func main() {
	minusC := flag.Bool("c", false, "Characters")
	minusW := flag.Bool("w", false, "Words")
	minusL := flag.Bool("l", false, "Lines")

	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Printf("usage: wc <file1> [<file2> [... <fileN]]\n")
		os.Exit(1)
	}

	totalLines := 0
	totalWords := 0
	totalCharacters := 0
	printAll := false

	for _, filename := range flag.Args() {
		numberOfLines, numberOfWords, numberOfCharacters := countLines(filename)

		totalLines = totalLines + numberOfLines
		totalWords = totalWords + numberOfWords
		totalCharacters = totalCharacters + numberOfCharacters

		if (*minusC && *minusW && *minusL) || (!*minusC && !*minusW && !*minusL) {
			fmt.Printf("%d", numberOfLines)
			fmt.Printf("\t%d", numberOfWords)
			fmt.Printf("\t%d", numberOfCharacters)
			fmt.Printf("\t%s\n", filename)
			printAll = true
			continue
		}

		if *minusL {
			fmt.Printf("%d", numberOfLines)
		}

		if *minusW {
			fmt.Printf("\t%d", numberOfWords)
		}

		if *minusC {
			fmt.Printf("\t%d", numberOfCharacters)
		}
		fmt.Printf("\t%s\n", filename)
	}

	if (len(flags) != 1) && printAll {
		fmt.Printf("%d", totalLines)
		fmt.Printf("\t%d", totalWords)
		fmt.Printf("\t%d", totalCharacters)
		fmt.Println("\ttotal")
		return
	}

	if (len(flags) != 1) && *minusL {
		fmt.Printf("%d", totalLines)
	}

	if (len(flags) != 1) && *minusW {
		fmt.Printf("\t%d", totalWords)
	}

	if (len(flags) != 1) && *minusC {
		fmt.Printf("\t%d", totalCharacters)
	}

	if len(flags) != 1 {
		fmt.Printf("\ttotal\n")
	}
}

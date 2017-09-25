package main

import (
	"fmt"
	"io"
	"os"
)

func Copy(src, dst string) (int64, error) {

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err

}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two command line arguments!")
		os.Exit(1)
	}

	sourceFile := os.Args[1]
	destinationFile := os.Args[2]
	nBytes, err := Copy(sourceFile, destinationFile)

	if err != nil {
		fmt.Printf("The copy operation failed %q\n", err)
	} else {
		fmt.Printf("Copied %d bytes!\n", nBytes)
	}
}

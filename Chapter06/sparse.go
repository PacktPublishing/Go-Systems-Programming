package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s SIZE filename\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	SIZE, _ := strconv.ParseInt(os.Args[1], 10, 64)
	filename := os.Args[2]

	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("File %s already exists.\n", filename)
		os.Exit(1)
	}

	fd, err := os.Create(filename)
	if err != nil {
		log.Fatal("Failed to create output")
	}

	_, err = fd.Seek(SIZE-1, 0)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to seek")
	}

	_, err = fd.Write([]byte{0})
	if err != nil {
		fmt.Println(err)
		log.Fatal("Write operation failed")
	}

	err = fd.Close()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to close file")
	}
}

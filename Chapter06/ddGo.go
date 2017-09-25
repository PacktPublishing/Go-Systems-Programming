package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createBytes(buf *[]byte, count int) {
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		intByte := byte(random(0, 9))
		*buf = append(*buf, intByte)
	}
}

func main() {
	minusBS := flag.Int("bs", 0, "Block Size")
	minusCOUNT := flag.Int("count", 0, "Counter")
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(-1)
	}

	if *minusBS < 0 || *minusCOUNT < 0 {
		fmt.Println("Count or/and Byte Size < 0!")
		os.Exit(-1)
	}

	filename := flags[0]
	rand.Seed(time.Now().Unix())

	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("File %s already exists.\n", filename)
		os.Exit(1)
	}

	destination, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create:", err)
		os.Exit(1)
	}

	buf := make([]byte, *minusBS)
	buf = nil
	for i := 0; i < *minusCOUNT; i++ {
		createBytes(&buf, *minusBS)
		if _, err := destination.Write(buf); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		buf = nil
	}
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Printf("usage: %s number\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}
	numGR, _ := strconv.ParseInt(os.Args[1], 10, 64)
	var waitGroup sync.WaitGroup
	var i int64

	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Printf("%d ", i)
		}()
	}
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}

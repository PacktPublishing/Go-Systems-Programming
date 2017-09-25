package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s integer\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	numGR, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Printf("Going to create %d goroutines.\n", numGR)
	var waitGroup sync.WaitGroup

	var i int64
	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(x int64) {
			defer waitGroup.Done()
			fmt.Printf(" %d ", x)
		}(i)
	}
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}

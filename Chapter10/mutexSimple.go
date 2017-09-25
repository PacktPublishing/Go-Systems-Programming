package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var aMutex sync.Mutex
var sharedVariable string = ""

func addDot() {
	aMutex.Lock()
	sharedVariable = sharedVariable + "."
	aMutex.Unlock()
}

func read() string {
	aMutex.Lock()
	a := sharedVariable
	aMutex.Unlock()
	return a
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s n\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	numGR, _ := strconv.ParseInt(os.Args[1], 10, 64)
	var waitGroup sync.WaitGroup

	var i int64
	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			addDot()
		}()
	}
	waitGroup.Wait()
	fmt.Printf("-> %s\n", read())
	fmt.Printf("Length: %d\n", len(read()))
}

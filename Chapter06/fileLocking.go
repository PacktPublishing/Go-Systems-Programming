package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var mu sync.Mutex

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func writeDataToFile(i int, file *os.File, w *sync.WaitGroup) {
	mu.Lock()
	time.Sleep(time.Duration(random(10, 1000)) * time.Millisecond)
	fmt.Fprintf(file, "From %d, writing %d\n", i, 2*i)
	fmt.Printf("Wrote from %d\n", i)
	w.Done()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide one command line argument!")
		os.Exit(-1)
	}

	filename := os.Args[1]
	number := 3

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var w *sync.WaitGroup = new(sync.WaitGroup)
	w.Add(number)

	for r := 0; r < number; r++ {
		go writeDataToFile(r, file, w)
	}

	w.Wait()
}

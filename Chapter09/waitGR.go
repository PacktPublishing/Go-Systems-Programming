package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Waiting for Goroutines!")

	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	var i int64
	for i = 0; i < 10; i++ {

		go func(x int64) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("\nExiting...")
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func SetValue(newValue int) {
	writeValue <- newValue
}

func ReadValue() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	go monitor()
	var waitGroup sync.WaitGroup

	for r := 0; r < 20; r++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			SetValue(rand.Intn(100))
		}()
	}
	waitGroup.Wait()
	fmt.Printf("\nLast value: %d\n", ReadValue())
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func addIntegers(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum = sum + input
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

func sendIntegers(c chan int) {
	for {
		c <- rand.Intn(100)
	}
}

func main() {
	c := make(chan int)
	go addIntegers(c)
	go sendIntegers(c)

	time.Sleep(2 * time.Second)
}

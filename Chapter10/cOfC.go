package main

import (
	"fmt"
)

var numbers = []int{0, -1, 2, 3, -4, 5, 6, -7, 8, 9, 10}

func f1(cc chan chan int, finished chan struct{}) {
	c := make(chan int)
	cc <- c
	defer close(c)

	total := 0
	i := 0
	for {
		select {
		case c <- numbers[i]:
			i = i + 1
			i = i % len(numbers)
			total = total + 1
		case <-finished:
			c <- total
			return
		}
	}
}

func main() {
	c1 := make(chan chan int)
	f := make(chan struct{})

	go f1(c1, f)
	numbers := <-c1

	i := 0
	for integer := range numbers {
		fmt.Printf("%d ", integer)
		i = i + 1
		if i == 100 {
			close(f)
		}
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"time"
)

func A(a, b, c chan struct{}) {
	<-a
	fmt.Println("A!")
	time.Sleep(time.Second)
	close(b)
}

func B(b, c chan struct{}) {
	<-b
	fmt.Println("B!")
	close(c)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C!")
}

func main() {
	a := make(chan struct{})
	b := make(chan struct{})
	c := make(chan struct{})

	go A(a, b, c)
	go C(c)
	go B(b, c)
	go C(c)

	close(a)
	time.Sleep(2 * time.Second)
}

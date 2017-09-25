package main

import (
	"fmt"
)

func change(x []int) {
	x[3] = -2
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}

func main() {
	aSlice := []int{-1, 4, 5, 0, 7, 9}
	fmt.Printf("Before change: ")
	printSlice(aSlice)
	change(aSlice)
	fmt.Printf("After change: ")
	printSlice(aSlice)

	fmt.Printf("Before. Cap: %d, length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("After. Cap: %d, length: %d\n", cap(aSlice), len(aSlice))
	printSlice(aSlice)
	anotherSlice := make([]int, 4)
	fmt.Printf("A new slice with 4 elements: ")
	printSlice(anotherSlice)
}

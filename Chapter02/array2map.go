package main

import (
	"fmt"
	"strconv"
)

func main() {
	anArray := [4]int{1, -2, 14, 0}
	aMap := make(map[string]int)

	length := len(anArray)
	for i := 0; i < length; i++ {
		fmt.Printf("%s ", strconv.Itoa(i))
		aMap[strconv.Itoa(i)] = anArray[i]
	}
	fmt.Printf("\n")

	for key, value := range aMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}

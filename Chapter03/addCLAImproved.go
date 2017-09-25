package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	sum := 0
	for i := 1; i < len(arguments); i++ {
		temp, err := strconv.Atoi(arguments[i])
		if err == nil {
			sum = sum + temp
		} else {
			fmt.Println("Ignoring", arguments[i])
		}
	}
	fmt.Println("Sum:", sum)
}

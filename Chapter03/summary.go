package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s [3]string
	s[0] = "1 b 3"
	s[1] = "11 a 1 14 1 1"
	s[2] = "-1 2 -3 -4 -5"

	arguments := os.Args
	column, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Error reading argument")
		os.Exit(-1)
	}
	if column == 0 {
		fmt.Println("Invalid column")
		os.Exit(1)
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		if len(data) >= column {
			temp, err := strconv.Atoi(data[column-1])
			if err == nil {
				sum = sum + temp
			} else {
				fmt.Printf("Invalid argument: %s\n", data[column-1])
			}
		} else {
			fmt.Println("Invalid column!")
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

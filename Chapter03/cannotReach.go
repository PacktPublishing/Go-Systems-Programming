package main

import (
	"fmt"
)

func x() int {
	return -1
	fmt.Println("Exiting x()")
}

func y() int {
	return -1
	fmt.Println("Exiting y()")
}

func main() {
	fmt.Println(x())
	fmt.Println("Exiting program...")
}

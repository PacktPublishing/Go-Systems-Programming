package main

import (
	"anotherPackage"
	"fmt"
)

func main() {
	anotherPackage.Version()
	// fmt.Println(anotherPackage.version)
	fmt.Println(anotherPackage.Pi)
}

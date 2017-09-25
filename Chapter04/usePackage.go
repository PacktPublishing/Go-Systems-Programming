package main

import (
	"aSimplePackage"
	"fmt"
)

func main() {
	temp := aSimplePackage.Add(5, 10)
	fmt.Println(temp)
	fmt.Println(aSimplePackage.Pi)
}

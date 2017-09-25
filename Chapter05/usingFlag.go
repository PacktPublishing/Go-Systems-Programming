package main

import (
	"flag"
	"fmt"
)

func main() {
	minusO := flag.Bool("o", false, "o")
	minusC := flag.Bool("c", false, "c")
	minusK := flag.Int("k", 0, "an int")

	flag.Parse()

	fmt.Println("-o:", *minusO)
	fmt.Println("-c:", *minusC)
	fmt.Println("-K:", *minusK)

	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}

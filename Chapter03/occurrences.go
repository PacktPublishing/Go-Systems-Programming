package main

import (
	"fmt"
	"strings"
)

func main() {

	var s [3]string
	s[0] = "1 b 3 1 a a b"
	s[1] = "11 a 1 1 1 1 a a"
	s[2] = "-1 b 1 -4 a 1"

	counts := make(map[string]int)

	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		for _, word := range data {
			_, ok := counts[word]
			if ok {
				counts[word] = counts[word] + 1
			} else {
				counts[word] = 1
			}
		}
	}

	for key, _ := range counts {
		fmt.Printf("%s -> %d \n", key, counts[key])
	}
}

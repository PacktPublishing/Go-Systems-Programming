package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 0
	TOTAL := 0
	if len(os.Args) > 3 {
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX, _ = strconv.Atoi(os.Args[2])
		TOTAL, _ = strconv.Atoi(os.Args[3])
	} else {
		fmt.Println("Usage:", os.Args[0], "MIX MAX TOTAL")
		os.Exit(-1)
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < TOTAL; i++ {
		myrand := random(MIN, MAX)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}

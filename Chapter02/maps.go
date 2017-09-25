package main

import (
	"fmt"
)

func main() {

	aMap := make(map[string]int)
	aMap["Mon"] = 0
	aMap["Tue"] = 1
	aMap["Wed"] = 2
	aMap["Thu"] = 3
	aMap["Fri"] = 4
	aMap["Sat"] = 5
	aMap["Sun"] = 6

	fmt.Printf("Sunday is the %dth day of the week.\n", aMap["Sun"])

	// Check if a key exists on a Map
	_, ok := aMap["Tuesday"]
	if ok {
		fmt.Printf("The Tuesday key exists!\n")
	} else {
		fmt.Printf("The Tuesday key does not exist!\n")
	}

	count := 0
	// Visit all Map elements
	for key, _ := range aMap {
		count++
		fmt.Printf("%s ", key)
	}
	fmt.Printf("\n")
	fmt.Printf("The aMap has %d elements\n", count)

	count = 0
	delete(aMap, "Fri")
	for _, _ = range aMap {
		count++
	}
	fmt.Printf("The aMap has now %d elements\n", count)

	anotherMap := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
	}
	anotherMap["Five"] = 5
	count = 0
	for _, _ = range anotherMap {
		count++
	}
	fmt.Printf("anotherMap has %d elements\n", count)
}

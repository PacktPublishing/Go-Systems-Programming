package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("Mihalis", "Mihalis Tsoukalos")
	fmt.Println(match)
	match, _ = regexp.MatchString("Tsoukalos", "Mihalis tsoukalos")
	fmt.Println(match)

	parse, err := regexp.Compile("[Mm]ihalis")

	if err != nil {
		fmt.Printf("Error compiling RE: %s\n", err)
	} else {
		fmt.Println(parse.MatchString("Mihalis Tsoukalos"))
		fmt.Println(parse.MatchString("mihalis Tsoukalos"))
		fmt.Println(parse.MatchString("M ihalis Tsoukalos"))
		fmt.Println(parse.ReplaceAllString("mihalis Mihalis", "MIHALIS"))
	}
}

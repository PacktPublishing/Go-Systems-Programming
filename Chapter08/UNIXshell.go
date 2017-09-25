package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var VERSION string = "0.2"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {

		line := scanner.Text()
		words := strings.Split(line, " ")
		command := words[0]

		switch command {
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		case "version":
			fmt.Println(VERSION)
		default:
			fmt.Println(line)
		}

		fmt.Print("> ")
	}
}

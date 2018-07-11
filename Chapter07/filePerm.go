package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func tripletToBinary(triplet string) string {
	if triplet == "rwx" {
		return "111"
	}
	if triplet == "-wx" {
		return "011"
	}
	if triplet == "--x" {
		return "001"
	}
	if triplet == "---" {
		return "000"
	}
	if triplet == "r-x" {
		return "101"
	}
	if triplet == "r--" {
		return "100"
	}
	if triplet == "--x" {
		return "001"
	}
	if triplet == "rw-" {
		return "110"
	}
	if triplet == "-w-" {
		return "010"
	}
	return "unknown"
}

func convertToBinary(permissions string) string {
	binaryPermissions := permissions[1:]
	p1 := binaryPermissions[0:3]
	p2 := binaryPermissions[3:6]
	p3 := binaryPermissions[6:9]
	return tripletToBinary(p1) + tripletToBinary(p2) + tripletToBinary(p3)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("usage: %s filename\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	filename := arguments[1]
	info, _ := os.Stat(filename)
	mode := info.Mode()

	fmt.Println(filename, "mode is", mode)
	fmt.Println("As string is", mode.String()[1:10])
	fmt.Println("As binary is", convertToBinary(mode.String()))
}

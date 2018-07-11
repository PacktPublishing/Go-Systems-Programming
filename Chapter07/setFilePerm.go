package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	p1 := permissions[0:3]
	p2 := permissions[3:6]
	p3 := permissions[6:9]

	p1 = tripletToBinary(p1)
	p2 = tripletToBinary(p2)
	p3 = tripletToBinary(p3)

	p1Int, _ := strconv.ParseInt(p1, 2, 64)
	p2Int, _ := strconv.ParseInt(p2, 2, 64)
	p3Int, _ := strconv.ParseInt(p3, 2, 64)

	returnValue := p1Int*100 + p2Int*10 + p3Int
	tempReturnValue := int(returnValue)
	returnString := "0" + strconv.Itoa(tempReturnValue)
	return returnString
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Printf("usage: %s filename permissions\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	filename, _ := filepath.EvalSymlinks(arguments[1])
	permissions := arguments[2]
	if len(permissions) != 9 {
		fmt.Println("Permissions should be 9 characters (rwxrwxrwx):", permissions)
		os.Exit(-1)
	}

	bin := convertToBinary(permissions)
	newPerms, _ := strconv.ParseUint(bin, 0, 32)
	newMode := os.FileMode(newPerms)
	os.Chmod(filename, newMode)
}

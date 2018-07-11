package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var PERMISSIONS string

func permissionsOfFIle(filename string) string {
	info, err := os.Stat(filename)
	if err != nil {
		return "-1"
	}
	mode := info.Mode()
	return mode.String()[1:10]
}

func walkFunction(path string, info os.FileInfo, err error) error {
	_, err = os.Lstat(path)
	if err != nil {
		return err
	}

	if permissionsOfFIle(path) == PERMISSIONS {
		fmt.Println(path)
	}
	return err
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Printf("usage: %s RootDirectory permissions\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	Path := arguments[1]
	Path, _ = filepath.EvalSymlinks(Path)
	PERMISSIONS = arguments[2]

	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
	}
}

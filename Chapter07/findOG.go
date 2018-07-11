package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("usage: %s <files>\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		fileInfo, err := os.Stat(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%+v\n", fileInfo.Sys())
		fmt.Println(fileInfo.Sys().(*syscall.Stat_t).Uid)
		fmt.Println(fileInfo.Sys().(*syscall.Stat_t).Gid)
	}
}

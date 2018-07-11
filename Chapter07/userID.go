package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		uid := os.Getuid()
		fmt.Println(uid)
		return
	}

	username := arguments[1]
	u, err := user.Lookup(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Uid)
}

package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	arguments := os.Args
	var u *user.User
	var err error
	if len(arguments) == 1 {
		u, err = user.Current()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		username := arguments[1]
		u, err = user.Lookup(username)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	gids, _ := u.GroupIds()
	for _, gid := range gids {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s(%s) ", group.Gid, group.Name)
	}
	fmt.Println()
}

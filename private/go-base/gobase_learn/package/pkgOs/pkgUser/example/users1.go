package main

import (
	"fmt"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usr.Gid)
	fmt.Println(usr.HomeDir)
	fmt.Println(usr.Name)
	fmt.Println(usr.Uid)
	fmt.Println(usr.Username)
	usr, _ = user.Lookup("antyiwei") //根据user name查找用户
	fmt.Println(usr)
	usr, err = user.LookupId("501") //根据userid查找用户
	fmt.Println(usr, err)
}

package main

import (
	"MONKE/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the MONKE Programming language! \n", user.Username)
	fmt.Printf("Type commands now!!")
	repl.Start(os.Stdin, os.Stdout)
}

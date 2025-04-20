package main

import (
	"fmt"
	"os"
	"os/user"
	"ta/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the ta Programming language! \n", user.Username)
	fmt.Printf("Type commands now!!")
	repl.Start(os.Stdin, os.Stdout)
}

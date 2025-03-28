package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! This is the Monkey programming language.\n", user.Username)
	fmt.Printf("Enter your commands:\n")
	startRepl(os.Stdin, os.Stdout)
}

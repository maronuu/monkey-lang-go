package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/maronuu/monkey-lang-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", user.Username)
	fmt.Print("Type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
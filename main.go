package main

import (
	"fmt"
	"monkey_interpreter/repl"
	"os"
	"os/user"
)

func main() {
	replUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", replUser.Username)
	fmt.Printf("Feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

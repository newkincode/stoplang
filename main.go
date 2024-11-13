package main

import (
	"fmt"
	"os"
	"os/user"
	"stop/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Stop v0.1\n",
		user.Username)
	fmt.Printf("\n")
	repl.Start(os.Stdin, os.Stdout)
}

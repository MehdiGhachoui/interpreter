package main

import (
	"fmt"
	"go-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This your programming laguage! \n", user.Username)
	fmt.Print("Type in commands to test\n")
	repl.Start(os.Stdin, os.Stdout)
}

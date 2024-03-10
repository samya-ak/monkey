package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/samya-ak/monkey/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to Monkey land! \n\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}

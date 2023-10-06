package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/JulienR1/monkey/internal/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	repl.Run(os.Stdin, os.Stdout)
}

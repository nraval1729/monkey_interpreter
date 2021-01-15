package main

import (
	"fmt"
	"./repl"
	"os"
)

func main() {
	fmt.Println("Welcome to the Monkey REPL!")
	fmt.Println("You know what to do, don't you?")

	repl.Start(os.Stdin, os.Stdout)
}
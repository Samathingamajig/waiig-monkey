package main

import (
	"fmt"
	"os"

	"github.com/Samathingamajig/waiig-monkey/repl"
)

func main() {
	fmt.Println("Hello! This is the Monkey programming language!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}

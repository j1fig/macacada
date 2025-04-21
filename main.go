package main

import (
	"fmt"
	"os"

	"monkey/repl"
)

const VERSION = "0.1"


func main() {
	fmt.Println("")
	fmt.Printf("  monkey v%s\n", VERSION)
	fmt.Println("")

	repl.Start(os.Stdin, os.Stdout)
}

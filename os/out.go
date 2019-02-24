package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "to stderr")
	fmt.Fprintln(os.Stdout, "to stdout")
}